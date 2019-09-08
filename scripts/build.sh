#!/usr/bin/env bash

set -e

function print_log() {
    echo "[$1] `date "+%Y-%m-%d %H:%M:%S"` --> $2"
}

print_log "INFO" "build begins..."
rm -f essai-api

print_log "INFO" "go build"
pwd
CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o essai-api ./main.go

print_log "INFO" "docker build"
docker build -t essai-api:latest .

#echo "---> docker push"
print_log "INFO" "docker push"
docker push aarondoge/essai-api:latest
