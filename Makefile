clear:
	rm -rf go_protos;

compile:
	mkdir -p go_protos;
	protoc \
		--proto_path=protos \
		protos/enums/**/*.proto \
		protos/messages/**/*.proto \
		--go_out=go_protos \
		--go_opt=paths=source_relative \
		--go-grpc_out=.;
