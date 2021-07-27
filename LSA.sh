#!/bin/zsh

if [ ! -d "./venv" ]; then
  python3 -m venv ./venv
fi
source venv/bin/activate
python3 nlp/LSA.py "$1" "$2"