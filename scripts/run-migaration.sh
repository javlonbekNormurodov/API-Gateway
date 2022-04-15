#!/usr/bin/env sh
set -e

>&2 echo "Running migration ..."
#migrate -path=./migrations -database=cassandra://127.0.0.1:9042/tutorialspoint?sslmode=disable
#migrate -path=./migrations -database=postgres://postgres:123@localhost:5432/postgres?sslmode=disable up
tail -f /dev/null