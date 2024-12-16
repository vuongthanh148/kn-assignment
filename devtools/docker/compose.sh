#!/usr/bin/env bash

# Script for running docker compose with the relevant compose.yaml and .env file.

source devtools/docker.sh || {
  echo "Are you at repo root?"
  exit 1
}

dockercompose $@
