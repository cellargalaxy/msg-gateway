#!/usr/bin/env bash

while :; do
  if [ ! -z "$configPath" ]; then
    break
  fi
  read -p "please enter config path(required):" configPath
done

if [ -z "$listenPort" ]; then
  read -p "please enter listen port(default:8990):" listenPort
fi
if [ -z "$listenPort" ]; then
  listenPort="8990"
fi

echo "configPath: $configPath"
echo "listenPort:$listenPort"
echo "input any key go on, or control+c over"
read

echo 'create volume'
docker volume create log
echo 'stop container'
docker stop msg_gateway
echo 'remove container'
docker rm msg_gateway
echo 'remove image'
docker rmi msg_gateway
echo 'docker build'
docker build -t msg_gateway .
echo 'docker run'
docker run -d \
  --restart=always \
  --name msg_gateway \
  -v log:/log \
  -p $listenPort:8990 \
  -v $configPath:/resources/config.ini \
  msg_gateway

echo 'all finish'
