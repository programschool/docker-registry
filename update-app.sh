#!/usr/bin/env bash


tag=`date +"TIME%Y-%m%d-%H%M"`

app="entry-proxy"
space="programschool-dev"

ID=`docker images registry.cn-wulanchabu.aliyuncs.com/$space/$app:latest -q`

docker pull registry.cn-wulanchabu.aliyuncs.com/$space/$app:latest

NEWID=`docker images registry.cn-wulanchabu.aliyuncs.com/$space/$app:latest -q`

if [[ $ID != $NEWID ]]
then
  execTag="docker tag $ID registry.cn-wulanchabu.aliyuncs.com/$space/$app:$tag"
  $execTag;
  stop="docker stop $app"
  $stop
  stop="docker rm $app"
  $rm

  docker-compose -f docker-compose.yml up -d
else
  echo -e "\n"
  echo no update
fi
