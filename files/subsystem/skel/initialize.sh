#!/bin/bash
# shellcheck disable=SC2164

cd "$HOME"/patent-server/

sudo apt install -y openjdk-8-jdk
sudo apt install -y python3.9
sudo apt install -y python3-pip
sudo apt install -y python3.9-venv
sudo apt install -y python3.9-dev

python3.9 -m venv venv
python3.9 -m pip install --upgrade pip
chmod 744 "$HOME"/patent-server/venv/bin/activate
source "$HOME"/patent-server/venv/bin/activate

pip3 install Cython
pip3 install wheel
pip3 install -r "$HOME"/patent-server/requirements.txt
