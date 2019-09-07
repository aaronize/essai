#!/usr/bin/env bash

set -e
. ./print.sh

deploy_type=$1

function development() {
    print_info_log "development deploy"
    docker run -d --rm --network=host --name essai-api -p 13030:13030 essai-api:latest # >/dev/null 2>&1
}

function production() {
    print_info_log "production deploy"
}

if [[ "$deploy_type" == "development" ]]; then
    development
elif [[ "$deploy_type" == "production" ]]; then
    production
else
    print_error_log  "invalid deploy type. support \"development/production\""
    exit 1
fi
