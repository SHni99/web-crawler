# go-web-crawler

A simple Go-based web crawler using gRPC. The server fetches a URL and returns all discovered links to the client.

## Prerequisites

- Go 1.23+ installed
- Protobuf compiler (`protoc`)
  ```sh
  brew install protobuf
  ```
- Go plugins:
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
- Ensure `$GOPATH/bin` (or `$(go env GOPATH)/bin`) is in your `PATH`.

## Generate gRPC code

```sh
protoc \
  --proto_path=proto \
  --go_out=proto --go_opt=paths=source_relative \
  --go-grpc_out=proto --go-grpc_opt=paths=source_relative \
  proto/crawler.proto

go mod tidy
```

## Build

```sh
# Build server
 go build -o bin/server ./cmd/server

# Build client
 go build -o bin/client ./cmd/client
```

## Run

1. Start server (default port 50051):
   ```sh
   ./bin/server -port 50051
   ```
2. In a new terminal, run client:
   ```sh
   ./bin/client -addr localhost:50051 -url <URL_TO_CRAWL> -timeout 10s
   ```

Replace `<URL_TO_CRAWL>` with the starting URL you wish to crawl.

## .gitignore

See `.gitignore` for ignored files and directories (binaries, logs, editor files, etc.).