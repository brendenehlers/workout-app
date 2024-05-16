#!/bin/bash

imageName=behlers22/wa-web-server
tag=latest

docker build \
  -f ./DOCKERFILE \
  -t $imageName \
  .

docker push \
  $imageName:$tag
