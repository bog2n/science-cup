.PHONY: build clean test

build:
	go build -ldflags="-w -s"

frontend/out:
	cd frontend && make

clean:
	rm -rf frontend/out

test:
	cd aminoacids && go test
	cd ribosome && go test
