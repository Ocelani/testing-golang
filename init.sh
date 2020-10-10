#!/bin/bash
#
# Initialize NestJS project for first usage.
#
# This script is following Google Shell Style standard
# https://google.github.io/styleguide/shell.xml

# STDERR log function
err() {
  echo -e "\n[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $@\n" >&2
  exit 1
}

# STDOUT log function
log() {
  echo -e "\n[$(date +'%Y-%m-%dT%H:%M:%S%z')]: $@\n"
}

log "yarn prebuild..."
yarn prebuild
if [ $? -ne 0 ]; then
    err "Error while yarn prebuild ..."
fi

log "yarn build..."
yarn build
if [ $? -ne 0 ]; then
    err "Error while yarn build ..."
fi

log "Starting Node production enviroment with yarn start:prod"
yarn start:prod
if [ $? -ne 0 ]; then
    err "Error while yarn start:prod ..."
fi