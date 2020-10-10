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

log "Golang is making unit testing ..."
gotest ./pkg -v
if [ $? -ne 0 ]; then
    err "Error while unit testing ..."
fi

log "Golang is making mutation testing ..."
go-mutesting ./pkg/main.go example/ github.com/zimmski/go-mutesting/mutator/...
if [ $? -ne 0 ]; then
    err "Error while mutation testing ..."
fi
