#!/usr/bin/env bash

set -e

exec_type=$1

function print_log() {
    echo "[$1] `date "+%Y-%m-%d %H:%M:%S"` --> $2"
}

function run() {
    docker run -d --rm --network=host --name essai-api -p 13030:13030 essai-api:latest # >/dev/null 2>&1
}

function stop() {
    docker stop essai-api # >/dev/null 2>&1
}

function restart() {
    stop; run
}

if [[ "$exec_type" == "run" ]]; then
    run
elif [[ "$exec_type" == "stop" ]]; then
    stop
elif [[ "$exec_type" == "restart" ]]; then
    restart
else
    print_log "ERR" "invalid exec type! support \"run/stop/restart\""
    exit 1
fi
