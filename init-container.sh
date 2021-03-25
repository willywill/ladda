#!/bin/bash

docker build -t wgermany/ladda-"${env:-test}":"${version:-1.0}" . &&
docker run -p "${port:-3001}":8080 --env-file ./.env wgermany/ladda-"${env:-test}":"${version:-1.0}"