#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { QuizAppStack } from '../lib/quiz-app-stack';

const app = new cdk.App();

new QuizAppStack(app, 'QuizAppStack', {
  // us-east-1 so a future CloudFront ACM certificate can live in this same
  // stack (CloudFront only accepts certificates from us-east-1).
  env: { region: 'us-east-1' },

  // Custom domain hook: when you own a Route 53 hosted zone, set these and
  // the stack wires up the ACM certificate, CloudFront alias and DNS record.
  // domainName: 'quiz.example.com',
  // hostedZoneDomain: 'example.com',
});
