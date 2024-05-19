#!/bin/bash

# config
ENV=dev

# vars
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# env-specfic config
/$SCRIPT_DIR/$ENV/setup.sh

# provision logging
ECK_VERSION=2.12.1
kubectl create -f https://download.elastic.co/downloads/eck/$ECK_VERSION/crds.yaml
kubectl apply -f https://download.elastic.co/downloads/eck/$ECK_VERSION/operator.yaml
kubectl apply -f $SCRIPT_DIR/logging/config/
kubectl apply -Rf $SCRIPT_DIR/logging/

# provision application
kubectl apply -f $SCRIPT_DIR/$ENV/config/
kubectl apply -Rf $SCRIPT_DIR/$ENV/
