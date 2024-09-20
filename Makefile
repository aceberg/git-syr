PKG_NAME=git-syr
DUSER=aceberg

mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/git-syr && \
	go mod tidy

run:
	cd cmd/git-syr/ && \
	go run . -c /data/git-syr/config.yaml -r /data/git-syr/repos.yaml -l /data/git-syr/git-syr.log

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

check: fmt lint

go-build:
	cd cmd/git-syr/ && \
	CGO_ENABLED=0 go build -o ../../tmp/git-syr .

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .

cli-run:
	cd cmd/git-syr-cli/ && \
	go run .

cli-build:
	cd cmd/git-syr-cli/ && \
	CGO_ENABLED=0 go build -o ../../tmp/git-syr-cli .