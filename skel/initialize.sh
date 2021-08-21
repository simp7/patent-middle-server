#!/bin/bash

# shellcheck disable=SC2164
cd "$HOME"/patent-server/
sudo apt install -y python3.8
sudo apt install -y python3-pip
sudo apt install -y python3.8-venv
sudo apt install -y python3.8-dev
python3.8 -m venv venv
chmod 744 "$HOME"/patent-server/venv/bin/activate
source "$HOME"/patent-server/venv/bin/activate
pip3 install Cython
pip3 install wheel

pip3 install -r "$HOME"/patent-server/requirements.txt
