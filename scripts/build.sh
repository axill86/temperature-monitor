#!/bin/bashd
docker run --rm -v $(pwd):/src -w /src golang:1.15 go -o out/ build -v