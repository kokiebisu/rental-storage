# rent-a-locker

# Onboarding (MacOS)

1 Create gmail account

2 Install Docker Desktop

3 Get invited to AWS organization

4 Create IAM user

5 Make sure you have NodeJS installed locally with version 16. web/storybook requires version 16.

6 Set up AWS PROFILE
`export AWS_PROFILE=rental-storage`

7 Get invited to Slack

8 Clone the repo

9 `brew install terraform golangci-lint flake8 mockery bash jq`

15 copy/paste ./.terraform.tfvars

16 `make setup` from root

# Onboarding (Windows)

1 Create gmail account

2 Install Docker Desktop

3 Get invited to AWS organization

4 Create IAM user

5 Set up AWS PROFILE
`setx AWS_PROFILE rental-storage`

6 Get invited to Slack

7 Clone the repo

8 `choco install terraform golangci-lint flake8 serverless`

9 Setting → Privacy & Security → For Developers → Turn on Developer Mode

10 Turn on Developer Mode
Search “Windows Features” on the search tab

11 Enable WSL on Windows
Open Turn windows feature on or off → Check Windows Subsystem for Linux → Click OK

12 Install Git Bash (This should be already installed when installing Git)

13 Install Ubuntu From Microsoft Store and setup Ubuntu (username, password)

14 I recommend to use Git Bash for the following commands

15 copy/paste ./.terraform.tfvars

16 `make setup` from root

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
