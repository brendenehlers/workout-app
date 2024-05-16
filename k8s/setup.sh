#!/bin/bash

# config
ENV=dev

# vars
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# configure ECK
ECK_VERSION=2.12.1

kubectl create -f https://download.elastic.co/downloads/eck/$ECK_VERSION/crds.yaml
kubectl apply -f https://download.elastic.co/downloads/eck/$ECK_VERSION/operator.yaml

# apply env
kubectl apply -Rf $SCRIPT_DIR/$ENV/
