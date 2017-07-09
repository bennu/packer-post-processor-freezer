#!/bin/bash
#
# build for multiple platforms.
set -e

OS=(linux windows darwin)
command_exists() {
  command -v "$@" > /dev/null 2>&1
}

bin() {
    printf ' Building freezer for %s...\n' $1

    mkdir -p bin/$1

    if [[ $1 == "windows" ]]; then
        suffix='.exe'
    fi

    docker run --rm -v $(pwd):/go/src/github.com/bennu/post-processor-freezer \
           -w /go/src/github.com/bennu/post-processor-freezer \
           -e GOARCH=amd64 -e GOOS=$1 -e CGO_ENABLED=0 golang:1.8-alpine  \
           go build -o bin/$1/packer-post-processor-freezer$suffix
}

if command_exists docker; then
    if [ $# -gt 0 ]; then
        bin $1
    else
    for os in "${OS[@]}"; do
        bin $os &
    done
    fi
    wait
else
  printf ' Error: you need to have Docker installed.'
fi