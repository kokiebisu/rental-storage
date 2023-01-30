# rent-a-locker

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

9 `brew install golangci-lint`

10 `brew install flake8`

11 `brew install mockery`

12 Must install latest bash for Mac users
`brew install bash`

13 `cd terraform && terraform init`

14 copy/paste ./.terraform.tfvars

15 `pnpm run terraform:setup:local` from root

16 `terraform apply`

17 Enter 'local', 'rental-storage'

18 `brew install serverless`

19 `pnpm run deps:install`

20 `pnpm run deploy:local`

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
