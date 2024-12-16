#!/usr/bin/env -S bash -e

# Script for dropping and creating a new database locally using docker.

docker rm -f local-postgres
./devtools/docker_postgres.sh
sleep 3 # wait for DB to be ready
go run ./devtools/cmd/db drop
./devtools/create_local_db.sh
