#!/usr/bin/env bash

function print_info_log() {
    echo "[INFO] `date "+%Y-%m-%d %H:%M:%S"` --> $1"
}

function print_error_log() {
    echo "[ERROR] `date "+%Y-%m-%d %H:%M:%S"` --> $1"
}
