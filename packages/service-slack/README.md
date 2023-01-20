# Slack Service

## Setup

### Prerequisites

- Installed homebrew (for Mac)

### Steps

#### Set up Virtual Environment

1. Install Pyenv
   On my Mac, I install pyenv & its sister project pyenv-virtualenv with Homebrew:

$ brew install readline xz
$ brew install pyenv pyenv-virtualenv

2. Make sure Pyenv is activated on every boot up
   You then need to add this to .bashrc (or .zshrc):

eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"
On Ubuntu, use the pyenv-installer:

$ sudo apt-get install -y make build-essential libssl-dev zlib1g-dev libbz2-dev \
 libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev libncursesw5-dev \
 xz-utils tk-dev libffi-dev liblzma-dev python-openssl git
$ curl https://pyenv.run | bash
You then need to add this to .bashrc:

$ export PATH="$HOME/.pyenv/bin:$PATH"
$ eval "$(pyenv init -)"
$ eval "$(pyenv virtualenv-init -)"
After restarting the terminal, pyenv will be available to you.

3. Install Python using Pyenv
   Pyenv’s main job is to install different python versions into their own environments and allow you to swap between them, You can even set it up so that it will try multiple versions in order when you run a Python application which can be quite useful.

To list the available versions: pyenv install -l. I install the latest versions of the Pythons that I’m interested in:

$ pyenv install 2.7.16
$ pyenv install 3.7.4
Use pyenv versions to see what’s installed.

We can now set a given version as our system-wide python with pyenv global, however, it’s much more useful to set up isolated environments and use them.

4. Create day-to-day environments
   Separate environments, known as virtualenvs or venvs, isolate an app and its dependencies from another one. In principle, you could have a separate environment for each application, but in practice, I’ve found that for my day-to-day apps, I can use the same environment for all apps for a given major Python version. I calls these environments apps2 and apps3 and put all my day-to-day apps and their dependencies in here, leaving the original Python installations clean for creating further environments for development work.

We create a new environment using the pyenv virtualenv command:

$ pyenv virtualenv 2.7.16 apps2
$ pyenv virtualenv 3.7.4 apps3
We set these are my system-wide defaults using pyenv global:

$ pyenv global apps3 apps2
This tells pyenv to look for a given app in the apps3 environment first and if it’s not there, look in apps2. We can now install python apps as required.

Firstly, the released version of rst2pdf:

#### Install dependencies

Make sure you are in the virtual environment and install the packages

`pip install -r requirements.txt`
