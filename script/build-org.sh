#!/bin/bash

cp ./conf.d/config-org.json config.json

build="docker build . -f Dockerfile -t org-apps.programschool.com/docker-registry:latest"
$build

push="docker push org-apps.programschool.com/docker-registry:latest"
$push

rm registry
