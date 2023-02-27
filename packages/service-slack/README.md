# Slack Service

## Setup

### Prerequisites

- Installed homebrew (for Mac)

### Steps

#### Set up Virtual Environment

`which python3` (Mac)
/usr/bin/python or /usr/bin/python3
-> there is something wrong if nothing is displayed

`python -V`
Python 2.7.12
-> If there is something wrong, nothing will be displayed

`brew install openssl readline sqlite3 xz zlib`

### Install Pyenv

What is pyenv?
Pyenv’s main job is to install different python versions into their own environments and allow you to swap between them, You can even set it up so that it will try multiple versions in order when you run a Python application which can be quite useful.

Using the pyenv-installer

`curl https://pyenv.run | bash`

If you are using Mac
`brew install pyenv`

### Add to the end of .zshrc/.bashrc

```
export PATH="$HOME/.pyenv/bin:$PATH"
eval "$(pyenv init -)"
```

Restart your terminal

### Install python version

`pyenv install 3.9.15`

### Check if the python version is properly installed

`ls ~/.pyenv/versions`
3.9.15

#### When you want to uninstall simply remove the folder

`pyenv uninstall 3.9.15`

#### Create new virtual environment

Separate environments, known as virtualenvs or venvs, isolate an app and its dependencies from another one. In principle, you could have a separate environment for each application, but in practice, I’ve found that for my day-to-day apps, I can use the same environment for all apps for a given major Python version. I calls these environments apps2 and apps3 and put all my day-to-day apps and their dependencies in here, leaving the original Python installations clean for creating further environments for development work.

`pyenv virtualenv <python_version> <environment_name>`
`pyenv virtualenv 3.9.15 slack-service`

#### Activate virtual environment

`pyenv activate slack-service`

##### Deactivate virtual environment

`pyenv deactivate`

### Install dependencies

Make sure you are in the virtual environment and install the packages

`pip install -r requirements.txt`

We can now set a given version as our system-wide python with pyenv global, however, it’s much more useful to set up isolated environments and use them.
