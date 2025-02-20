clear:
	rm -rf go_protos

compile:
	protoc \
		--proto_path=protos protos/**/*.proto \
		--go_out=. \
		--go-grpc_out=.
