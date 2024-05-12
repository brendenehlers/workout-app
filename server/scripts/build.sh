#!/bin/bash

CGO_ENABLED=0
GOOS=linux

rm -r dist
mkdir -p dist

go build \
  -a \
  -installsuffix cgo \
  -o dist/app \
  ./cmd/app/.