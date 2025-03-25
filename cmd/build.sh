#!/usr/bin/env bash

go build -ldflags "-s -w" -buildmode=c-shared -o grdp.so main.go
