export GOBIN=${PWD}/bin
GO=go
GOFMT=gofmt


all: build

build:
	$GOFMT -w . && $GO install  ./... 

clean:
	$GO clean -i ./...

test:
	$GO test -race ./...

.PHONY: storage
storage:
	$GO install -a github.com/flyaways/storage/cmd/agent



