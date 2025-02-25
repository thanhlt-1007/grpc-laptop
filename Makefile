clear:
	rm -rf go_protos;

compile:
	mkdir -p go_protos;
	protoc \
		--proto_path=protos \
		protos/enums/**/*.proto \
		protos/messages/**/*.proto \
		protos/services/**/*.proto \
		--go_out=go_protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=go_protos \
		--go-grpc_opt=paths=source_relative \

test:
	go clean -testcache;
	go test ./tests/...

run_server:
	go run cmd/server/main.go -port 8080
