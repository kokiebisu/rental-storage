# rent-a-locker

# stack

- go (no domain services (authorizer, authentication, payment, relational database services (domains), nosql database services (transactions))
- python (data processing (numpy, machine learning eventually?), notification (slack sdk), email sendgrid)
- serverless framework
- cloudformation
- stripe
- aws (lambda, sqs, sns, s3, cloudformation, kinesis, cloudwatch. rds. dynamodb, systems manager)
- postgres

# Add query/mutation

1 Add query to mapping-template/datasource/schema.graphql (api)
2 Add query/mutation to serverless.yml in service package
3 Add code

# Onboarding

1 Create gmail account

2 Install Docker Desktop

3 Get invited to AWS organization

4 Create IAM user

5 Set up AWS PROFILE
`export AWS_PROFILE=rental-storage`
`setx AWS_PROFILE rental-storage`

6 Get invited to Slack

7 Clone the repo

8 `brew install terraform`

9 Must install latest bash for Mac users
`brew install bash`

9 `cd terraform && terraform init`

10 copy/paste ./.terraform.tfvars

11 `pnpm run terraform:setup:local` from root

12 `terraform apply`

13 Enter 'local', 'rental-storage'

14 `brew install serverless`

15 `pnpm run deps:install`

16 `pnpm run deploy:local`

## Enable Git Hooks

1 You can enable python linting when committing your changes automatically by using the following command.
`make enable-precommit`

You can disable the pre-commit hook by the following
`make disable-precommit`

# terraform

local
use local state

dev/staging
use s3 backend (github actions)
