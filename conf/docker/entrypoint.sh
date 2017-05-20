#!/usr/bin/env bash
cd ${GOPATH}/src/github.com/spaiz/gogameoflife/tests && go test -v
cd ${GOPATH}/src/github.com/spaiz/gogameoflife/cmd/golife && go install
golife "$@"