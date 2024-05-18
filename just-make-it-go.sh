#!/bin/bash

# this'll take the service from nothing to working.

minikube start

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
/$SCRIPT_DIR/k8s/setup.sh

kubectl port-forward -n ingress-nginx service/ingress-nginx-controller 8080:80