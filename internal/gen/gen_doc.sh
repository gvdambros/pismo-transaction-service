#!/bin/bash

if [ -z "$BASH" ]; then echo "Please run this script with bash"; exit 1; fi

SCRIPT_PATH=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_PATH="$(cd $SCRIPT_PATH && cd ../../ && pwd)"
DOCS_PATH="$SCRIPT_PATH/doc"

mkdir -p "$DOCS_PATH"

cd "$ROOT_PATH"
swag --version
swag init --pd -g cmd/main.go --output "$DOCS_PATH"