#!/usr/bin/env -S bash -e
# Script for creating a new database locally.

source devtools/lib.sh || {
  echo "Are you at repo root?"
  exit 1
}

go run devtools/cmd/db/main.go create
./devtools/migrate_db.sh up
