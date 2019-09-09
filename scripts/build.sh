#!/usr/bin/env bash

set -e

CGO_ENABLED=0

function print_log() {
    echo "[$1] `date "+%Y-%m-%d %H:%M:%S"` --> $2"
}

print_log "INFO" "build begins..."
rm -f essai-api

print_log "INFO" "go build"
#GO_BUILD -o essai-api ./main.go
docker run --rm -v /data/jenkins_home/workspace/essai-api:/go/src/essai-api golang:1.12-alpine \
pwd && go build -a -ldflags '-extldflags "-static"' -o /go/src/essai-apiessai-api /go/src/essai-api/main.go

print_log "INFO" "docker build"
docker build -t essai-api:latest ../

#echo "---> docker push"
print_log "INFO" "docker push"
docker push aarondoge/essai-api:latest
