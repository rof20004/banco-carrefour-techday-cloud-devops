#!/bin/sh

# Esse script é utilizado no Dockerfile para inicializar os dois processos no container Docker: aplicação em go e o nginx

# Start api
nohup ./backend > backend.log 2>&1 &

# Start nginx
nginx -g 'daemon off;'