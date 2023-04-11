#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_FOLDER=$(cd "$SCRIPT_DIR" && cd ../../../../ && pwd)

cd "$ROOT_FOLDER" && \
docker build --no-cache -t test-postgres:latest \
-f "$SCRIPT_DIR"/Dockerfile .