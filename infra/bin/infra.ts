#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { QuizAppStack } from '../lib/quiz-app-stack';

const app = new cdk.App();

// Custom domain: pass via `-c siteDomainNames=a,b -c apiDomainName=... -c
// certificateArn=... -c cloudflareOriginSecret=...` (or set them in
// cdk.json's "context"). All four must be set together, or the stack falls
// back to the default CloudFront/API Gateway domains. See README "Dominio
// propio (Cloudflare)".
const siteDomainNamesCtx = app.node.tryGetContext('siteDomainNames') as string | undefined;
const apiDomainName = app.node.tryGetContext('apiDomainName') as string | undefined;
const certificateArn = app.node.tryGetContext('certificateArn') as string | undefined;
const cloudflareOriginSecret = app.node.tryGetContext('cloudflareOriginSecret') as string | undefined;

// Cost alert: pass via `-c budgetAlertEmail=you@example.com`. Independent of
// the custom-domain context above — see QuizAppStackProps.budgetAlertEmail.
const budgetAlertEmail = app.node.tryGetContext('budgetAlertEmail') as string | undefined;

new QuizAppStack(app, 'QuizAppStack', {
  // us-east-1 so a CloudFront ACM certificate can live in this same stack
  // (CloudFront only accepts certificates from us-east-1).
  env: { region: 'us-east-1' },

  siteDomainNames: siteDomainNamesCtx?.split(','),
  apiDomainName,
  certificateArn,
  cloudflareOriginSecret,
  budgetAlertEmail,
});
