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

2 Get invited to AWS organization

3 Create IAM user

4 Set up AWS PROFILE
`export AWS_PROFILE=rental-storage`
`setx AWS_PROFILE rental-storage`

5 Get invited to Slack

6 Clone the repo

7 Install root packages

8 brew install terraform

9 cd terraform && terraform init

10 create .terraform.tfvars

11 replace profile with 'rental-storage'
