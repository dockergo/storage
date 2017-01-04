export GOBIN=${PWD}/docker/bin

all: build

build:
	gofmt -w . && go install  ./... 

clean:
	go clean -i ./...

test:
	go test -race ./...


.PHONY: agent
agent:
	go install -a github.com/flyaways/storage/cmd/agent



