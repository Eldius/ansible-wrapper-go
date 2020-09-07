#!/bin/bash

eval "$(pyenv init -)" || exit 1
eval "$(pyenv virtualenv-init -)" || exit 1

ansible --version && \
    echo "Executing script..." || \
        exit 1

#sleep 10