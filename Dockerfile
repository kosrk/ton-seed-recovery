FROM docker.io/library/golang:1.22-bullseye as builder

WORKDIR /go/src/github.com/kosrk/ton-seed-recovery/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd cmd

RUN go build -o bin/recovery ./cmd/

FROM ubuntu:20.04 as runner
RUN apt-get update && \
    apt-get install -y openssl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/github.com/kosrk/ton-seed-recovery/bin/recovery .

ENTRYPOINT ["/recovery"]
