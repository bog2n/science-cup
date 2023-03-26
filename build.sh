#!/bin/sh

set -e

# Build frontend

docker run --rm -v "$PWD"/frontend:/app -w /app alpine:3.15 \
	sh -c 'apk add nodejs npm make; make'

# Build releases

mkdir -p releases

GOOS=linux
for GOARCH in amd64 arm arm64; do
	export GOOS GOARCH
	make build
	mv science-cup releases/dnanalyzer-${GOOS}-${GOARCH}
done

GOOS=darwin
for GOARCH in amd64 arm64; do
	export GOOS GOARCH
	make build
	mv science-cup releases/dnanalyzer-${GOOS}-${GOARCH}
done

GOOS=windows
GOARCH=amd64
export GOOS GOARCH
make build
mv science-cup releases/dnanalyzer-${GOOS}-${GOARCH}.exe
