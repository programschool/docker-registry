#!/bin/bash

export GOOS=linux
export GOARCH=amd64

go mod vendor
go fmt main.go
go build -o registry main.go

if [[ $1 = 'local' ]]
then
  bash script/build-local.sh
elif [[ $1 = 'online' ]]
then
  bash script/build-online.sh "$2"
elif [[ $1 = 'org' ]]
then
  bash script/build-org.sh "$2"
fi
