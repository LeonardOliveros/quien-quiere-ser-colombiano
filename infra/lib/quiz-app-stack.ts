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
import * as acm from 'aws-cdk-lib/aws-certificatemanager';
import * as budgets from 'aws-cdk-lib/aws-budgets';

const repoRoot = path.join(__dirname, '..', '..');

export interface QuizAppStackProps extends cdk.StackProps {
  /**
   * Custom domain, all optional together. DNS lives in Cloudflare (not
   * Route53), so the certificate is requested out-of-band (AWS CLI) and
   * validated manually via a Cloudflare CNAME, then imported here by ARN —
   * see README "Dominio propio (Cloudflare)". Without these, the stack
   * behaves as before: CloudFront/API Gateway default domains only.
   */
  readonly siteDomainNames?: string[];
  readonly apiDomainName?: string;
  readonly certificateArn?: string;
  /**
   * Shared secret Cloudflare injects as the `X-Origin-Verify` header (via a
   * Transform Rule) on every request it forwards. CloudFront (viewer-request
   * function) and the Lambda both reject requests missing/mismatching it, so
   * direct hits to the CloudFront/API Gateway default domains — bypassing
   * Cloudflare's proxy/DDoS protection — get a 403 instead of reaching the
   * origin. Required alongside the other three domain props; see README.
   */
  readonly cloudflareOriginSecret?: string;
  /**
   * Email to alert when actual monthly AWS cost crosses budgetAlertLimitUsd
   * (see Costos in README) — nothing in this stack rate-limits abusive
   * traffic per IP, so this is the last-resort tripwire against a runaway
   * bill from the pay-per-use resources (Lambda, DynamoDB on-demand).
   * Optional: without it, no CfnBudget is created.
   */
  readonly budgetAlertEmail?: string;
}

const BUDGET_ALERT_LIMIT_USD = 10;

export class QuizAppStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: QuizAppStackProps) {
    super(scope, id, props);

    const domainsConfigured = !!(
      props?.siteDomainNames?.length &&
      props?.apiDomainName &&
      props?.certificateArn &&
      props?.cloudflareOriginSecret
    );
    const certificate = props?.certificateArn
      ? acm.Certificate.fromCertificateArn(this, 'Certificate', props.certificateArn)
      : undefined;

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
        ...(props?.siteDomainNames?.length
          ? { ALLOWED_ORIGINS: props.siteDomainNames.map((d) => `https://${d}`).join(',') }
          : {}),
        // Only set when Cloudflare fronts a custom domain (see
        // cloudflareOriginSecret above) — checked by originVerifyRequired()
        // in main.go so direct calls to the API Gateway default/custom
        // domain (bypassing Cloudflare) get a 403.
        ...(domainsConfigured ? { ORIGIN_VERIFY_SECRET: props!.cloudflareOriginSecret! } : {}),
      },
    });
    table.grantReadWriteData(apiFn);

    // Billing visibility: applied to every taggable resource in the stack.
    cdk.Tags.of(this).add('Project', 'QuienQuiereSerColombiano');
    cdk.Tags.of(this).add('Repository', 'https://github.com/LeonardOliveros/quien-quiere-ser-colombiano');

    // Regional custom domain for the API (api.<domain>), separate from the
    // site's CloudFront distribution — Cloudflare CNAMEs straight to it.
    const apiDomain = domainsConfigured
      ? new apigwv2.DomainName(this, 'ApiDomainName', {
          domainName: props!.apiDomainName!,
          certificate: certificate!,
        })
      : undefined;

    const httpApi = new apigwv2.HttpApi(this, 'HttpApi', {
      defaultIntegration: new HttpLambdaIntegration('ApiIntegration', apiFn, {
        payloadFormatVersion: apigwv2.PayloadFormatVersion.VERSION_2_0,
      }),
      ...(apiDomain ? { defaultDomainMapping: { domainName: apiDomain } } : {}),
    });

    // ------------------------------------------------------------------- web
    const siteBucket = new s3.Bucket(this, 'SiteBucket', {
      blockPublicAccess: s3.BlockPublicAccess.BLOCK_ALL,
      encryption: s3.BucketEncryption.S3_MANAGED,
      enforceSSL: true,
      removalPolicy: cdk.RemovalPolicy.DESTROY, // only holds a copy of dist/
      autoDeleteObjects: true,
    });

    // Static media (background music, future audio/video assets), separate
    // from siteBucket because that one is fully replaced (prune: true) on
    // every frontend deploy — a media file would survive that here. Served
    // through the same CloudFront distribution under /media/* (see
    // additionalBehaviors below) so the frontend can reference it with a
    // same-origin relative path, no CORS/env var needed.
    const mediaBucket = new s3.Bucket(this, 'MediaBucket', {
      blockPublicAccess: s3.BlockPublicAccess.BLOCK_ALL,
      encryption: s3.BucketEncryption.S3_MANAGED,
      enforceSSL: true,
      removalPolicy: cdk.RemovalPolicy.DESTROY, // reproducible from infra/assets/media
      autoDeleteObjects: true,
    });

    // Viewer-request function scoped to the S3 behavior only. It does two
    // things:
    //  1. When Cloudflare is configured (domainsConfigured), reject any
    //     request missing the `X-Origin-Verify` secret header — this is what
    //     stops someone from bypassing Cloudflare's proxy/DDoS protection by
    //     calling the CloudFront default domain directly. Cloudflare adds
    //     the header via a Transform Rule (see README); it can't be a
    //     distribution-wide WAF rule because CloudFront Functions are the
    //     cheapest way to do this without provisioning a WAF WebACL.
    //  2. SPA fallback: extension-less paths are Vue Router routes and must
    //     serve index.html. Distribution-wide errorResponses would also
    //     rewrite legitimate API 403/404 JSON bodies into index.html, so this
    //     is scoped to the function instead.
    const originVerifyCheck = domainsConfigured
      ? `
  var secretHeader = headers['x-origin-verify'];
  if (!secretHeader || secretHeader.value !== ${JSON.stringify(props!.cloudflareOriginSecret)}) {
    return {
      statusCode: 403,
      statusDescription: 'Forbidden',
      body: { encoding: 'text', data: 'Forbidden' },
    };
  }`
      : '';
    const spaRewriteFn = new cloudfront.Function(this, 'SpaRewrite', {
      comment: 'Verify Cloudflare origin secret; rewrite extension-less URIs to /index.html (SPA fallback)',
      code: cloudfront.FunctionCode.fromInline(`
function handler(event) {
  var request = event.request;
  var headers = request.headers;
${originVerifyCheck}
  if (!request.uri.split('/').pop().includes('.')) {
    request.uri = '/index.html';
  }
  return request;
}
`),
    });

    // With a custom domain, the API lives on its own subdomain (api.<domain>,
    // via apiDomain above) and the frontend is built with an absolute
    // VITE_API_BASE_URL, so no /api/* behavior is needed here. Without one
    // (default AWS domains), the frontend still calls the relative '/api',
    // so this distribution keeps proxying it to API Gateway same-origin.
    const distribution = new cloudfront.Distribution(this, 'Distribution', {
      comment: 'Quiz app: SPA from S3' + (domainsConfigured ? '' : ', /api/* to API Gateway'),
      defaultRootObject: 'index.html',
      ...(domainsConfigured ? { domainNames: props!.siteDomainNames, certificate } : {}),
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
        // Long-lived edge cache: media/ objects are named with a version
        // suffix (e.g. himno-nacional-instrumental-v1.mp3) instead of being
        // overwritten in place, so CACHING_OPTIMIZED's ~24h default TTL never
        // serves stale bytes — a new version just gets a new URL.
        '/media/*': {
          origin: origins.S3BucketOrigin.withOriginAccessControl(mediaBucket),
          viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
          cachePolicy: cloudfront.CachePolicy.CACHING_OPTIMIZED,
          // Reuse the same function as the default behavior so the
          // Cloudflare origin-verify check (bypass protection) also covers
          // media requests; its SPA-rewrite half is a no-op here since every
          // media path has a file extension.
          functionAssociations: [{
            function: spaRewriteFn,
            eventType: cloudfront.FunctionEventType.VIEWER_REQUEST,
          }],
        },
        ...(domainsConfigured ? {} : {
          '/api/*': {
            // apiEndpoint is https://<id>.execute-api.<region>.amazonaws.com
            origin: new origins.HttpOrigin(cdk.Fn.select(2, cdk.Fn.split('/', httpApi.apiEndpoint))),
            viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
            allowedMethods: cloudfront.AllowedMethods.ALLOW_ALL,
            cachePolicy: cloudfront.CachePolicy.CACHING_DISABLED,
            // Forward everything except Host: API Gateway routes on its own host.
            originRequestPolicy: cloudfront.OriginRequestPolicy.ALL_VIEWER_EXCEPT_HOST_HEADER,
          },
        }),
      },
    });

    new s3deploy.BucketDeployment(this, 'DeploySite', {
      sources: [s3deploy.Source.asset(path.join(repoRoot, 'dist'))],
      destinationBucket: siteBucket,
      distribution, // invalidate the CDN on every deploy
      distributionPaths: ['/*'],
      prune: true,
    });

    // mediaBucket's contents are NOT managed by CDK/BucketDeployment: media
    // files (e.g. background music) are uploaded directly with `aws s3 cp`
    // (see README "Media assets (S3 + CloudFront)") instead of being
    // committed to the repo as a CDK asset. Object key must start with
    // `media/` to match the /media/* behavior above (no originPath
    // stripping is configured), e.g.:
    //   aws s3 cp <file> s3://<MediaBucketName output>/media/audio/<name>.mp3
    //   aws cloudfront create-invalidation --distribution-id <id> --paths "/media/audio/<name>.mp3"

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

    // ----------------------------------------------------------- cost alert
    // Account-wide (not scoped to this stack's resources) since tag-based
    // budget filters need cost-allocation tags activated manually in the
    // Billing console first — an account-wide cap is simpler and reliable
    // out of the box, which fits an AWS account used mainly for this project.
    if (props?.budgetAlertEmail) {
      new budgets.CfnBudget(this, 'CostAlertBudget', {
        budget: {
          budgetType: 'COST',
          timeUnit: 'MONTHLY',
          budgetLimit: { amount: BUDGET_ALERT_LIMIT_USD, unit: 'USD' },
        },
        notificationsWithSubscribers: [
          {
            notification: {
              notificationType: 'ACTUAL',
              comparisonOperator: 'GREATER_THAN',
              threshold: 100,
              thresholdType: 'PERCENTAGE',
            },
            subscribers: [{ subscriptionType: 'EMAIL', address: props.budgetAlertEmail }],
          },
        ],
      });
    }

    // --------------------------------------------------------------- outputs
    new cdk.CfnOutput(this, 'DistributionUrl', { value: `https://${distribution.distributionDomainName}` });
    new cdk.CfnOutput(this, 'ApiEndpoint', { value: httpApi.apiEndpoint });
    new cdk.CfnOutput(this, 'TableName', { value: table.tableName });
    new cdk.CfnOutput(this, 'FunctionName', { value: apiFn.functionName });
    new cdk.CfnOutput(this, 'MediaBucketName', { value: mediaBucket.bucketName });
    if (apiDomain) {
      // CNAME target for api.<domain> in Cloudflare (proxied).
      new cdk.CfnOutput(this, 'ApiCustomDomainTarget', { value: apiDomain.regionalDomainName });
    }
  }
}
