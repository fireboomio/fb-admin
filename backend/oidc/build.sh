#!/usr/bin/env bash
export GOPROXY="https://goproxy.cn,direct"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o casdoor_service .