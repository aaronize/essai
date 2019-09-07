#!/usr/bin/env bash

set -e

deploy_type=$1

function print_info_log() {
    echo "[INFO] `date "+%Y-%m-%d %H:%M:%S"` --> $1"
}

function print_error_log() {
    echo "[ERROR] `date "+%Y-%m-%d %H:%M:%S"` --> $1"
}

function development() {
    print_info_log "development deploy begin..."
    docker run -d --rm --network=host --name essai-api -p 13030:13030 essai-api:latest # >/dev/null 2>&1
}

function production() {
    print_info_log "production deploy begin..."
}

if [[ "$deploy_type" == "development" ]]; then
    development
elif [[ "$deploy_type" == "production" ]]; then
    production
else
    print_error_log  "invalid deploy type. support \"development/production\""
    exit 1
fi
