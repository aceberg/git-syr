FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/git-syr/ && CGO_ENABLED=0 go build -o /git-syr .


FROM alpine

WORKDIR /app

RUN apk add --no-cache git openssh-client ca-certificates bash \
    && mkdir -p /data/git-syr \
    && mkdir -p /root/.ssh

COPY --from=builder /git-syr /app/
COPY config /root/.ssh/

ENTRYPOINT ["./git-syr"]