#!/bin/bash

eval "$(pyenv init -)" || exit 1
eval "$(pyenv virtualenv-init -)" || exit 1

ansible-playbook --version 

#sleep 10