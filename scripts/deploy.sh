#!/usr/bin/env bash

set -e

deploy_type=$1

function print_log() {
    echo "[$1] `date "+%Y-%m-%d %H:%M:%S"` --> $2"
}

function development() {
    print_log "INFO" "development deploy"
    docker run -d --rm --network=host --name essai-api -p 13030:13030 essai-api:latest # >/dev/null 2>&1
}

function production() {
    print_log "INFO" "production deploy"
}

if [[ "$deploy_type" == "development" ]]; then
    development
elif [[ "$deploy_type" == "production" ]]; then
    production
else
    print_log "ERR" "invalid deploy type. support \"development/production\""
    exit 1
fi
