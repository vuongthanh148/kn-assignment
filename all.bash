#!/usr/bin/env bash

source devtools/lib.sh || {
  echo "Are you at repo root?"
  exit 1
}

GO=go

# Support ** in globs, for check_script_hashes.
shopt -s globstar

warnout() {
  while read line; do
    warn "$line"
  done
}

# filter FILES GLOB1 GLOB2 ...
# returns the files in FILES that match any of the glob patterns.
filter() {
  local infiles=$1
  shift
  local patterns=$*
  local outfiles=

  for pat in $patterns; do
    for f in $infiles; do
      if [[ $f == $pat ]]; then
        outfiles="$outfiles $f"
      fi
    done
  done
  echo $outfiles
}

# Return the files that are modified or added.
# If there are such files in the working directory, whether or not
# they are staged for commit, use those.
# Otherwise, use the files changed since the previous commit.
modified_files() {
  local working=$(working_files)
  if [[ $working != ' ' ]]; then
    echo $working
  elif [[ $(git rev-parse HEAD) = $(git rev-parse master) ]]; then
    echo ""
  else
    diff_files HEAD^
  fi
}

# Return the files in the working directory that have not been added to the commit.
working_files() {
  echo "$(diff_files '') $(diff_files --cached)"
}

# Helper for modified_files. It asks git for all modified, added or deleted
# files, and keeps only the latter two.
diff_files() {
  git diff --name-status $* | awk '$1 ~ /^R/ { print $3; next } $1 != "D" { print $2 }'
}

# ensure_go_binary verifies that a binary exists in $PATH corresponding to the
# given go-gettable URI. If no such binary exists, it is fetched via `go install`.
ensure_go_binary() {
  local binary=$(basename $1)
  if ! [ -x "$(command -v $binary)" ]; then
    info "Installing: $1"
    # Run in a subshell for convenience, so that we don't have to worry about
    # our PWD.
    (
      set -x
      cd && $GO install $1@latest
    )
  fi
}

# bad_migrations outputs migrations with bad sequence numbers.
bad_migrations() {
  ls migrations | cut -d _ -f 1 | sort | uniq -c | grep -vE '^\s+2 '
}

# check_bad_migrations looks for sql migration files with bad sequence numbers,
# possibly resulting from a bad merge.
check_bad_migrations() {
  info "Checking for bad migrations"
  bad_migrations | while read line; do
    err "unexpected number of migrations: $line"
  done
}

# check_vet runs go vet on source files.
check_vet() {
  runcmd $GO vet -all ./...
}

# check_staticcheck runs staticcheck on source files.
check_staticcheck() {
  ensure_go_binary honnef.co/go/tools/cmd/staticcheck
  runcmd staticcheck $(go list ./... | grep -v third_party | grep -v internal/doc | grep -v internal/render)
}

# check_misspell runs misspell on source files.
check_misspell() {
  ensure_go_binary github.com/client9/misspell/cmd/misspell
  runcmd misspell cmd/**/*.{go,sh} internal/**/* README.md
}

go_linters() {
  check_vet
  check_staticcheck
  check_misspell
}

standard_linters() {
  check_bad_migrations
  go_linters
}

usage() {
  cat <<EOUSAGE
Usage: $0 [subcommand]
Available subcommands:
  help           - display this help message
  (empty)        - run all standard checks and tests
  ci             - run checks and tests suitable for continuous integration
  cl             - run checks and tests on the current CL, suitable for a commit or pre-push hook
  lint           - run all standard linters below:
  migrations     - (lint) check migration sequence numbers
  misspell       - (lint) run misspell on source files
  staticcheck    - (lint) run staticcheck on source files
EOUSAGE
}

# Packages to run without the race detector on CI.
# (They time out with -race.)
declare -A no_race
no_race=(
)

main() {
  case "$1" in
  "-h" | "--help" | "help")
    usage
    exit 0
    ;;
  "")
    standard_linters
    runcmd $GO mod tidy
    runcmd env GO_DISCOVERY_TESTDB=true go test ./...
    runcmd $GO test ./internal/secrets
    ;;
  cl)
    # Similar to the above, but only run checks that apply to files in this commit.
    files=$(modified_files)
    if [[ $files = '' ]]; then
      info "No modified files; nothing to do."
      exit 0
    fi
    info "Running checks on:"
    info "    " $files

    if [[ $(filter "$files" 'migrations/*') != '' ]]; then
      check_bad_migrations
    fi
    if [[ $(filter "$files" '*.go') != '' ]]; then
      go_linters
    fi
    runcmd $GO mod tidy
    runcmd env GO_DISCOVERY_TESTDB=true go test ./...
    runcmd $GO test ./internal/secrets
    ;;

  ci)
    # Similar to the no-arg mode, but omit actions that require GCP
    # permissions or that don't test the code.
    # Also, run the race detector on most tests.
    local start=$(date +%s)

    # Explicitly mark the working directory as safe in CI.
    # https://github.com/docker-library/golang/issues/452
    local wd=$(pwd)
    runcmd git config --system --add safe.directory ${wd}

    standard_linters
    # Print how long it takes to download dependencies and run the standard
    # linters in CI.
    local end=$(date +%s)
    echo
    echo "--------------------"
    echo "DONE: $((end - start)) seconds"
    echo "--------------------"

    for pkg in $($GO list ./...); do
      if [[ ${no_race[$pkg]} = '' ]]; then
        race="$race $pkg"
      fi
    done
    runcmd env GO_DISCOVERY_TESTDB=true $GO test -race -count=1 $race
    runcmd env GO_DISCOVERY_TESTDB=true $GO test -count=1 ${!no_race[*]}
    ;;
  lint) standard_linters ;;
  migrations) check_migrations ;;
  misspell) check_misspell ;;
  staticcheck) check_staticcheck ;;

  *)
    usage
    exit 1
    ;;
  esac
  if [[ $EXIT_CODE != 0 ]]; then
    err "FAILED; see errors above"
  fi
  exit $EXIT_CODE
}

main $@
