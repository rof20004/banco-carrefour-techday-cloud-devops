#!/bin/sh

# Start api
nohup ./backend > backend.log 2>&1 &

# Start nginx
nginx -g 'daemon off;'