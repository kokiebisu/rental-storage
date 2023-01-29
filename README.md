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

11 Must install latest bash for Mac users
`brew install bash`

12 `cd terraform && terraform init`

13 copy/paste ./.terraform.tfvars

14 `pnpm run terraform:setup:local` from root

15 `terraform apply`

16 Enter 'local', 'rental-storage'

17 `brew install serverless`

18 `pnpm run deps:install`

19 `pnpm run deploy:local`

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
