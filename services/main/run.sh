#!/usr/bin/env bash
HOST_URI='https://wallet.traaittxtcash.com' \
HOST_PORT=':443' \
USER_URI='http://localhost:8081' \
WALLET_URI='http://localhost:8082' \
/usr/local/go/bin/go run main.go init.go handlers.go utils.go
