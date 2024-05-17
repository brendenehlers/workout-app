#!/bin/bash


imageName=behlers22/wa-static-files
tag=latest

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd $SCRIPT_DIR

docker build \
  -f $SCRIPT_DIR/DOCKERFILE \
  -t $imageName \
  .

docker push \
  $imageName:$tag