# rent-a-locker

# Onboarding (MacOS)

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

13 Install latest bash
`brew install bash`

14 `cd terraform && terraform init`

15 copy/paste ./.terraform.tfvars

16 `pnpm run terraform:setup:local` from root

17 `terraform apply`

18 Enter 'local', 'rental-storage'

19 `brew install serverless`

20 `pnpm run deps:install`

21 `pnpm run deploy:local`

# Onboarding (Windows)

1 Create gmail account

2 Install Docker Desktop

3 Get invited to AWS organization

4 Create IAM user

5 Set up AWS PROFILE
`setx AWS_PROFILE rental-storage`

6 Get invited to Slack

7 Clone the repo

8 `choco install terraform`

9 `choco install golangci-lint`

10 `choco install flake8`

11 Setting → Privacy & Security → For Developers → Turn on Developer Mode

12 Turn on Developer Mode
Search “Windows Features” on the search tab

13 Enable WSL on Windows
Open Turn windows feature on or off → Check Windows Subsystem for Linux → Click OK

14 Install Git Bash (This should be already installed when installing Git)

15 Install Ubuntu From Microsoft Store and setup Ubuntu (username, password)

16 I recommend to use Git Bash for the following commands

17 `cd terraform && terraform init`

18 copy/paste ./.terraform.tfvars

19 `pnpm run terraform:setup:local` from root

20 `terraform apply`

21 Enter 'local', 'rental-storage'

22 `choco install serverless`

23 `pnpm run deps:install`

24 `pnpm run deploy:local`

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
