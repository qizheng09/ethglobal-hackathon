#!/usr/bin/env bash

mkdir -p output/bin output/config
cp script/* output/
cp config/*.toml output/config/
chmod +x output/boostrap.sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/bin/hackathon

tar -czvf hackathon.tar.gz output/*
#scp hackathon.tar.gz root@47.252.32.182:/root/hackathon/