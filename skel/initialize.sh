#!/bin/bash

# shellcheck disable=SC2164
cd "$HOME"/patent-server/
sudo apt install -y python3
sudo apt install -y python3-pip
sudo apt install -y python3-venv
python3 -m venv venv
chmod 744 "$HOME"/patent-server/venv/bin/activate
source "$HOME"/patent-server/venv/bin/activate
pip3 install Cython
pip3 install -r "$HOME"/patent-server/requirements.txt
