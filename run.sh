#!/bin/bash

export BLACK_MARLIN_API_CONSUL_URL="localhost:8500"
export BLACK_MARLIN_API_CONSUL_PATH="black_marlin"
export BLACK_MARLIN_SYSTEM_KEY="31561f28-2a8e-447c-87e5-2fac01ab6bbb"
export BLACK_MARLIN_SERVICE_KEY="31561f28-2a8e-447c-87e5-2fac01ab6eee"

GOARCH="amd64"
GOOS="darwin"
# GOOS="linux"
CGO_ENABLED="0"

cmd=$1

binary="black-marlin-api"

glide install
go build -v -o ${binary}
#./${binary} serve
./${binary} migration auto
