#!/bin/bash

# This is used for locally running the application in a container
# docker build -t wgermany/ladda-"${env:-test}":"${version:-1.0}" . &&
# docker run -v "${path:-$(dirname $(readlink -f $0))/temp}":/app/files -p "${port:-3001}":8080 --env-file ./.env wgermany/ladda-"${env:-test}":"${version:-1.0}"

# This is used for running the application in a container on the web server
docker pull ghcr.io/willywill/ladda:latest &&
docker run -v "${path:-($PWD)/temp}":/app/files -p "${port:-3001}":3001 --env-file ./.env ghcr.io/willywill/ladda:latest