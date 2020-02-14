#!/bin/bash

GOOS=js GOARCH=wasm go build -o main.wasm
command -v goexec >/dev/null || go get -u github.com/shurcooL/goexec
echo "Starting ..."
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
