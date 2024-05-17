#!/bin/bash

# DEV SETUP

minikube addons enable metrics-server
minikube addons enable ingress

echo "Waiting for ingress to start..."
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=120s
echo "Ingress started!"