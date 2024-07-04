#!/usr/bin/env bash
DB_USER=root \
DB_PWD=password \
HOST_URI='http://localhost' \
HOST_PORT=':8081' \
WALLET_URI='http://localhost:8082' go run users.go utils.go
