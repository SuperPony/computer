#!/bin/bash

sudo wget https://github.com/jesseduffield/lazydocker/releases/download/v0.12/lazydocker_0.12_Linux_x86_64.tar.gz

sudo mkdir ./lazydocker \
     && mv ./lazydocker*.tar.gz ./lazydocker \
     && tar -zxvf lazydocker*.tar.gz  ./lazydocker