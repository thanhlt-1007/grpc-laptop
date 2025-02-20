clear:
	rm -rf go-protos

compile:
	mkdir -p go-protos;
	protoc \
		--proto_path=protos protos/**/*.proto \
		--go_out=go-protos --go_opt=paths=source_relative \
		--go-grpc_out=.
