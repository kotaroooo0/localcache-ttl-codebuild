FROM golang:1.14-alpine as builder

WORKDIR /go/src/github.com/kotaroooo0/localcache-ttl-codebuild

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

# キャッシュが活用されるかわかりやすくするためにスリープを挟む
# RUN sleep 120

FROM alpine:latest

COPY --from=builder /go/src/github.com/kotaroooo0/localcache-ttl-codebuild /main

ENTRYPOINT ["/main"]
