#!/bin/bash

cp ./conf.d/config-local.json config.json

build="docker build . -f local/Dockerfile -t registry.com:5000/docker-registry:latest"
$build

push="docker push registry.com:5000/docker-registry:latest"
$push

rm registry
