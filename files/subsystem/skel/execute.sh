#!/bin/bash

export OPENBLAS_NUM_THREADS=1
source "$HOME"/patent-server/venv/bin/activate
python3.9 "$@"
