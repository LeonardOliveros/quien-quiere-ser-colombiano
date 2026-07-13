import * as path from 'path';
import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { GoFunction } from '@aws-cdk/aws-lambda-go-alpha';
import * as apigwv2 from 'aws-cdk-lib/aws-apigatewayv2';
import { HttpLambdaIntegration } from 'aws-cdk-lib/aws-apigatewayv2-integrations';
import * as s3 from 'aws-cdk-lib/aws-s3';
import * as s3deploy from 'aws-cdk-lib/aws-s3-deployment';
import * as cloudfront from 'aws-cdk-lib/aws-cloudfront';
import * as origins from 'aws-cdk-lib/aws-cloudfront-origins';
import * as cr from 'aws-cdk-lib/custom-resources';
import * as iam from 'aws-cdk-lib/aws-iam';

const repoRoot = path.join(__dirname, '..', '..');

export interface QuizAppStackProps extends cdk.StackProps {
  /**
   * Custom domain hook (not wired yet — deploys use the CloudFront URL).
   * When set together with hostedZoneDomain, add to this stack:
   *   1. route53.HostedZone.fromLookup({ domainName: hostedZoneDomain })
   *   2. new acm.Certificate({ domainName, validation: fromDns(zone) })
   *      (this stack already lives in us-east-1, as CloudFront requires)
   *   3. Distribution props: { domainNames: [domainName], certificate }
   *   4. new route53.ARecord({ zone, recordName: domainName,
   *        target: RecordTarget.fromAlias(new targets.CloudFrontTarget(distribution)) })
   */
  readonly domainName?: string;
  readonly hostedZoneDomain?: string;
}

export class QuizAppStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: QuizAppStackProps) {
    super(scope, id, props);

    // ------------------------------------------------------------------ data
    // Single table, generic PK/SK keys, no GSIs (see internal/storage/dynamodb).
    const table = new dynamodb.TableV2(this, 'Table', {
      partitionKey: { name: 'PK', type: dynamodb.AttributeType.STRING },
      sortKey: { name: 'SK', type: dynamodb.AttributeType.STRING },
      billing: dynamodb.Billing.onDemand(),
      timeToLiveAttribute: 'ttl',
      removalPolicy: cdk.RemovalPolicy.RETAIN, // keep user data on stack delete
    });

    // ------------------------------------------------------------------- api
    const apiFn = new GoFunction(this, 'ApiFunction', {
      entry: repoRoot, // main package at the repo root
      runtime: lambda.Runtime.PROVIDED_AL2023,
      architecture: lambda.Architecture.ARM_64,
      memorySize: 512,
      timeout: cdk.Duration.seconds(30),
      bundling: {
        environment: { CGO_ENABLED: '0' },
      },
      environment: {
        DB_DRIVER: 'dynamodb',
        DYNAMODB_TABLE: table.tableName,
        SEED_ON_START: 'false', // seeding runs at deploy time (custom resource)
        GIN_MODE: 'release',
      },
    });
    table.grantReadWriteData(apiFn);

    const httpApi = new apigwv2.HttpApi(this, 'HttpApi', {
      defaultIntegration: new HttpLambdaIntegration('ApiIntegration', apiFn, {
        payloadFormatVersion: apigwv2.PayloadFormatVersion.VERSION_2_0,
      }),
    });

    // ------------------------------------------------------------------- web
    const siteBucket = new s3.Bucket(this, 'SiteBucket', {
      blockPublicAccess: s3.BlockPublicAccess.BLOCK_ALL,
      encryption: s3.BucketEncryption.S3_MANAGED,
      enforceSSL: true,
      removalPolicy: cdk.RemovalPolicy.DESTROY, // only holds a copy of dist/
      autoDeleteObjects: true,
    });

    // SPA fallback scoped to the S3 behavior only: extension-less paths are
    // Vue Router routes and must serve index.html. Distribution-wide
    // errorResponses would also rewrite legitimate API 403/404 JSON bodies
    // into index.html, so a viewer-request function is used instead.
    const spaRewriteFn = new cloudfront.Function(this, 'SpaRewrite', {
      comment: 'Rewrite extension-less URIs to /index.html (SPA fallback)',
      code: cloudfront.FunctionCode.fromInline(`
function handler(event) {
  var request = event.request;
  if (!request.uri.split('/').pop().includes('.')) {
    request.uri = '/index.html';
  }
  return request;
}
`),
    });

    // API Gateway origin: apiEndpoint is https://<id>.execute-api.<region>.amazonaws.com
    const apiDomain = cdk.Fn.select(2, cdk.Fn.split('/', httpApi.apiEndpoint));

    const distribution = new cloudfront.Distribution(this, 'Distribution', {
      comment: 'Quiz app: SPA from S3, /api/* to API Gateway',
      defaultRootObject: 'index.html',
      defaultBehavior: {
        origin: origins.S3BucketOrigin.withOriginAccessControl(siteBucket),
        viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        cachePolicy: cloudfront.CachePolicy.CACHING_OPTIMIZED,
        functionAssociations: [{
          function: spaRewriteFn,
          eventType: cloudfront.FunctionEventType.VIEWER_REQUEST,
        }],
      },
      additionalBehaviors: {
        '/api/*': {
          origin: new origins.HttpOrigin(apiDomain),
          viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
          allowedMethods: cloudfront.AllowedMethods.ALLOW_ALL,
          cachePolicy: cloudfront.CachePolicy.CACHING_DISABLED,
          // Forward everything except Host: API Gateway routes on its own host.
          originRequestPolicy: cloudfront.OriginRequestPolicy.ALL_VIEWER_EXCEPT_HOST_HEADER,
        },
      },
    });

    new s3deploy.BucketDeployment(this, 'DeploySite', {
      sources: [s3deploy.Source.asset(path.join(repoRoot, 'dist'))],
      destinationBucket: siteBucket,
      distribution, // invalidate the CDN on every deploy
      distributionPaths: ['/*'],
      prune: true,
    });

    // ---------------------------------------------------------------- seeder
    // Deploy-time seeding: invoke the API function with the seed action. The
    // question bank ships inside the binary (go:embed), so the trigger
    // combines the data/ fingerprint with the function version — a new
    // question file or new code re-runs the seeder, nothing else does.
    const dataFingerprint = cdk.FileSystem.fingerprint(path.join(repoRoot, 'data'));
    const seedCall: cr.AwsSdkCall = {
      service: 'Lambda',
      action: 'invoke',
      parameters: {
        FunctionName: apiFn.functionName,
        Payload: JSON.stringify({ quizapp_action: 'seed' }),
      },
      physicalResourceId: cr.PhysicalResourceId.of(`seed-${dataFingerprint}-${apiFn.currentVersion.version}`),
    };
    const seeder = new cr.AwsCustomResource(this, 'Seeder', {
      onCreate: seedCall,
      onUpdate: seedCall,
      // fromSdkCalls infers the IAM action from the SDK method name ("invoke"
      // -> "lambda:Invoke"), which isn't a real IAM action — the actual one
      // is "lambda:InvokeFunction" — so it must be spelled out explicitly.
      policy: cr.AwsCustomResourcePolicy.fromStatements([
        new iam.PolicyStatement({
          actions: ['lambda:InvokeFunction'],
          resources: [apiFn.functionArn],
        }),
      ]),
      installLatestAwsSdk: false,
      timeout: cdk.Duration.minutes(2),
    });
    seeder.node.addDependency(table);

    // --------------------------------------------------------------- outputs
    new cdk.CfnOutput(this, 'DistributionUrl', { value: `https://${distribution.distributionDomainName}` });
    new cdk.CfnOutput(this, 'ApiEndpoint', { value: httpApi.apiEndpoint });
    new cdk.CfnOutput(this, 'TableName', { value: table.tableName });
    new cdk.CfnOutput(this, 'FunctionName', { value: apiFn.functionName });
  }
}
