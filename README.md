# rent-a-locker

# Onboarding (macOS)

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

11 Install latest bash
`brew install bash`

12 `cd terraform && terraform init`

13 copy/paste ./.terraform.tfvars

14 `pnpm run terraform:setup:local` from root

15 `terraform apply`

16 Enter 'local', 'rental-storage'

17 `brew install serverless`

18 `pnpm run deps:install`

19 `pnpm run deploy:local`

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
