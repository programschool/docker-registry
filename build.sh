#!/bin/bash

export GOOS=linux
export GOARCH=amd64

go fmt main.go
go build -o registry main.go

cp ./conf/config-dev.json config.json
dev="-dev"

if [[ $1 = '--prod' ]]
then
  cp ./conf/config.json config.json
  dev=""
fi

build="docker build . -f Dockerfile -t registry.cn-wulanchabu.aliyuncs.com/programschool$dev/registry:latest"
$build

push="docker push registry.cn-wulanchabu.aliyuncs.com/programschool$dev/registry:latest"
$push

rm registry
