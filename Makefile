mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/GitSyncTimer && \
	go mod tidy

run:
	cd cmd/GitSyncTimer/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run
	golint ./...

go-build:
	cd cmd/GitSyncTimer/ && \
	CGO_ENABLED=0 go build -o ../../tmp/GitSyncTimer .