clear:
	rm -rf go_protos

compile:
	protoc \
		--proto_path=protos protos/**/*.proto \
		--go_out=go-protos --go_opt=paths=source_relative \
		--go-grpc_out=.
