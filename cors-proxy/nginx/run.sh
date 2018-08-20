#!/usr/bin/env bash

docker run -e VIRTUAL_HOST=clickbait --rm --name dev-nginx -p 80:80 -v $PWD/nginx.conf:/etc/nginx/nginx.conf:ro  nginx:1.11.8-alpine