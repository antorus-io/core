#!/bin/sh

go test -cover -coverprofile ./coverage/coverage.out ./... > /dev/null

exec "$@"
