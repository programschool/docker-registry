#!/usr/bin/env bash

bash ./build.sh
scp main root@192.168.10.103:/home/services/registry
rm main
