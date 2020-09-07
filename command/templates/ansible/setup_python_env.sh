#!/bin/bash

eval "$(pyenv init -)" || exit 1
eval "$(pyenv virtualenv-init -)" || exit 1

pyenv install 3.8.0
pyenv virtualenv 3.8.0 ansible-wrapper
pyenv local 3.8.0/envs/ansible-wrapper
cd {{ .Workspace }}
pip install ansible
