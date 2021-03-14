#!/usr/bin/env bash


tag=`date +"TIME%Y-%m%d-%H%M"`

app="registry"
space="programschool-dev"

ID=`docker images registry.cn-wulanchabu.aliyuncs.com/$space/$app:latest -q`

docker pull registry.cn-wulanchabu.aliyuncs.com/$space/$app:latest

NEWID=`docker images registry.cn-wulanchabu.aliyuncs.com/$space/$app:latest -q`

if [[ $ID != $NEWID ]]
then
  execTag="docker tag $ID registry.cn-wulanchabu.aliyuncs.com/$space/$app:$tag"
  $execTag;
  stopcmd="docker stop $app"
  $stopcmd
  rmcmd="docker rm $app"
  $rmcmd

  /usr/local/bin/docker-compose -f docker-compose.yml up -d
else
  echo -e "\n"
  echo no update
fi
