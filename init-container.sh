#!/bin/bash

docker build -t wgermany/ladda-"${env:-test}":"${version:-1.0}" . &&
docker volume create --name LaddaVolume &&
docker run --rm -v LaddaVolume:/laddavolume -p "${port:-3001}":8080 --env-file ./.env wgermany/ladda-"${env:-test}":"${version:-1.0}"