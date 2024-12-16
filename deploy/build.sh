#!/usr/bin/env bash
set -e

source private/devtools/lib.sh || {
  echo "Are you at repo root?"
  exit 1
}

usage() {
  cat >&2 <<EOUSAGE

  Usage: $0 PROJECT_ID BUILD_TAG

  Build and push docker containers for the worker and frontend services.

EOUSAGE
  exit 1
}

main() {
  local project_id=$1
  local build_tag=$2
  for service in worker frontend; do
    runcmd docker build -t gcr.io/$project_id/$service:$build_tag \
      -f private/config/Dockerfile.$service .
    runcmd docker push gcr.io/$project_id/$service:$build_tag
  done
}

main $@
