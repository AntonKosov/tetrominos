#!/bin/bash

golangci-lint run || { echo 'lint failed' ; exit 1; }

rm ./build/*
env GOOS=windows GOARCH=amd64 go build -o ./build/tetrominos.windows.amd64.exe .
env GOOS=linux GOARCH=amd64 go build -o ./build/tetrominos.linux.amd64 .
env GOOS=darwin GOARCH=amd64 go build -o ./build/tetrominos.darwin.amd64 .