# rent-a-locker

# Onboarding

1 Create gmail account

2 Install Docker Desktop

3 Get invited to AWS organization

4 Create IAM user

5 Make sure you have NodeJS installed locally with version 16. web/storybook requires version 16.

6 Set up AWS PROFILE
`export AWS_PROFILE=rental-storage`
`setx AWS_PROFILE rental-storage`

7 Get invited to Slack

8 Clone the repo

9 `brew install terraform`

10 `brew install golangci-lint`

11 `brew install flake8`

12 `brew install mockery`

13 Must install latest bash for Mac users
`brew install bash`

14 `cd terraform && terraform init`

15 copy/paste ./.terraform.tfvars

16 `pnpm run terraform:setup:local` from root

17 `terraform apply`

18 Enter 'local', 'rental-storage'

19 `brew install serverless`

20 `pnpm run deps:install`

21 `pnpm run deploy:local`

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
