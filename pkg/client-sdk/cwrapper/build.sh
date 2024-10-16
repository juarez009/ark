#!/bin/bash

set -e

go build -buildmode=c-shared -o libark.so main.go
