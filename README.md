# grpc-laptop

- TECH SCHOOL gRPC

- Reference: https://www.youtube.com/watch?v=YzypniHHU3w&list=PLy_6D98if3UJd5hxWNfAqKMr15HZqFnqf&index=6

## gvm

```sh
gvm install go1.24.0
gvm use go1.24.0
```

## gofumpt

- Install

```sh
go install mvdan.cc/gofumpt@v0.7.0
```

- Run

```sh
gofumpt -l -w .
```

## golangci-lint

- Install

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
```

- Run

```sh
golangci-lint run
```

## protoc

```sh
sudo apt-get install -y protobuf-compiler
protoc --version # libprotoc 29.3
```

## protoc-gen-go

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5
protoc-gen-go --version # protoc-gen-go v1.36.5
```

## protoc-gen-go-grpc

```sh
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
protoc-gen-go-grpc --version # protoc-gen-go-grpc 1.5.1
```
