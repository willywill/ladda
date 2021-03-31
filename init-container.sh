#!/bin/bash

# This file is used for locally running the application in a container
docker build -t wgermany/ladda-"${env:-test}":"${version:-1.0}" . &&
docker run -v "${path:-$(dirname $(readlink -f $0))/temp}":/app/files -p "${port:-3001}":8080 --env-file ./.env wgermany/ladda-"${env:-test}":"${version:-1.0}"