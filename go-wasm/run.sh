#!/bin/bash
set -ex

if ! command -v goexec; then
    go get -u github.com/shurcooL/goexec
fi

goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
