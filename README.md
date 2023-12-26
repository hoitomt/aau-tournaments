# AAU Tournament Scraper

## Dev

### Setup

1. `go mod tidy`
1. Protocol Buffers
   1. Install protoc tool: `brew install protobuf`
   1. Install protobufs plugin `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

### Run

1. `go build`
2. Compile protobufs (If they have changed):

    ```
    protoc -I=protos --go_out=protos protos/*.proto
    ```

2. `./aau-tournaments`