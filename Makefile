.PHONY: build clean test
VERSION := $(shell git describe --tags --dirty)

docker:
	docker build . -t dnanalyzer:$(VERSION)

build:
	go build -ldflags="-w -s"

frontend/out:
	cd frontend && make

clean:
	rm -rf frontend/out

test:
	cd aminoacids && go test
	cd ribosome && go test
