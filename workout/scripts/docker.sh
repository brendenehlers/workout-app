#!/bin/bash

imageName=behlers22/wa-workout-service
tag=latest

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $SCRIPT_DIR/..

docker build \
  -f ./DOCKERFILE \
  -t $imageName \
  .

docker push \
  $imageName:$tag
