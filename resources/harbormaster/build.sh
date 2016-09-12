#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

mkdir -p /home/drydock/go/src/github.com/etcinit
ln -sf $(pwd) /home/drydock/go/src/github.com/etcinit/
cd /home/drydock/go/src/github.com/etcinit/gonduit

glide install
go build $(glide novendor)
go test $(glide novendor)
