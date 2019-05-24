#!/bin/bash

export CGO_ENABLED=0
export GOOS=windows
export GOARCH=amd64

sh -c 'cd check-cpu && go get && go build -ldflags "-s -w" -o ../dist/check-cpu.exe'
sh -c 'cd check-disk-usage && go get && go build -ldflags "-s -w" -o ../dist/check-disk-usage.exe'
sh -c 'cd check-memory-percent && go get && go build -ldflags "-s -w" -o ../dist/check-memory-percent.exe'
