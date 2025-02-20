clear:
	rm -rf go-protos;
	clear

compile:
	mkdir -p go-protos;
	protoc \
		--proto_path=protos \
		protos/enums/**/*.proto \
		protos/messages/*.proto \
		--go_out=go-protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=.
