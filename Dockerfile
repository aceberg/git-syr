FROM golang:alpine AS builder

RUN apk add build-base
COPY . /src
RUN cd /src/cmd/GitSyncTimer/ && CGO_ENABLED=0 go build -o /GitSyncTimer .


FROM alpine

WORKDIR /app

RUN apk add --no-cache git openssh-client ca-certificates bash \
    && mkdir -p /data/GitSyncTimer \
    && mkdir -p /root/.ssh

COPY --from=builder /GitSyncTimer /app/
COPY config /root/.ssh/

ENTRYPOINT ["./GitSyncTimer"]