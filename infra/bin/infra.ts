#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { QuizAppStack } from '../lib/quiz-app-stack';

const app = new cdk.App();

// Custom domain: pass via `-c siteDomainNames=a,b -c apiDomainName=... -c
// certificateArn=...` (or set them in cdk.json's "context"). All three must
// be set together, or the stack falls back to the default CloudFront/API
// Gateway domains. See README "Dominio propio (Cloudflare)".
const siteDomainNamesCtx = app.node.tryGetContext('siteDomainNames') as string | undefined;
const apiDomainName = app.node.tryGetContext('apiDomainName') as string | undefined;
const certificateArn = app.node.tryGetContext('certificateArn') as string | undefined;

new QuizAppStack(app, 'QuizAppStack', {
  // us-east-1 so a CloudFront ACM certificate can live in this same stack
  // (CloudFront only accepts certificates from us-east-1).
  env: { region: 'us-east-1' },

  siteDomainNames: siteDomainNamesCtx?.split(','),
  apiDomainName,
  certificateArn,
});
