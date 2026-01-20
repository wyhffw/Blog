#!/bin/bash
export DATA_DIR=./data/posts
export PUBLIC_DIR=./web/dist
export ADMIN_USER=admin
export ADMIN_PASS=test123
go run ./cmd/server
