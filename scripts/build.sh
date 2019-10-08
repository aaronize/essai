#!/usr/bin/env bash

set -e

# CGO_ENABLED=0     # 这个变量放在这里不起作用
build_type=$1

function print_log() {
    echo "[$1] `date "+%Y-%m-%d %H:%M:%S"` --> $2"
}

print_log "INFO" "build begins..."
function go_build() {
    rm -f essai-api

    print_log "INFO" "go build"
    CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o /go/src/essai/essai-api /go/src/essai/main.go
}

function docker_build() {
    print_log "INFO" "docker build"
    docker build -t aarondoge/essai-api:latest .

    #echo "---> docker push"
    print_log "INFO" "docker push"
#    docker push aarondoge/essai-api:latest
}

if [[ "$build_type" == "golang" ]]; then
    go_build
elif [[ "$build_type" == "docker" ]]; then
    docker_build
else
    print_log "ERROR" "invalid build type"
    exit 1
fi


# -------------------------

#set -e
#
#function print_log() {
#    echo "[$1] `date "+%Y-%m-%d %H:%M:%S"` --> $2"
#}
#
#print_log "INFO" "build begins..."
#rm -f essai-api
#
#print_log "INFO" "go build"
#CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o /go/src/essai/essai-api /go/src/essai/main.go
#
#print_log "INFO" "docker build"
#docker build -t essai-api:latest .
#
##echo "---> docker push"
#print_log "INFO" "docker push"
#docker push aarondoge/essai-api:latest
