PKG_NAME=git-syr
DUSER=aceberg

mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/git-syr && \
	go mod tidy

run:
	cd cmd/git-syr/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

go-build:
	cd cmd/git-syr/ && \
	CGO_ENABLED=0 go build -o ../../tmp/git-syr .

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .