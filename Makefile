mod:
	rm go.mod || true && \
	rm go.sum || true && \
	go mod init github.com/aceberg/GitBackup && \
	go mod tidy

run:
	cd cmd/GitBackup/ && \
	go run .

fmt:
	go fmt ./...

lint:
	golangci-lint run

go-build:
	cd cmd/GitBackup/ && \
	CGO_ENABLED=0 go build -o ../../GitBackup .