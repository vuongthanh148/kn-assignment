#!/usr/bin/env bash
set -e

usage() {
  cat >&2 <<EOUSAGE

  Usage: $0 PROJECT_ID

  Clone the private repo from Cloud Source repos, then generate a build tag
  for the deployment and an ID token for IAP requests.

EOUSAGE
  exit 1
}

main() {
  local project_id=$1
  gcloud source repos clone private private
  source private/devtools/lib.sh || {
    echo "Are you at repo root?"
    exit 1
  }
  (cd private && docker_image_tag >../_BUILD_TAG)
  private/devtools/idtoken.sh $project_id >_ID_TOKEN
}

main $@
