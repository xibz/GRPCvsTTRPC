generate: generate-grpc-proto generate-ttrpc-proto

generate-grpc-proto:
	protoc -I models/grpcmodels --go_out=plugins=grpc:models/grpcmodels models/grpcmodels/*.proto

generate-ttrpc-proto:
	protoc -I models/ttrpcmodels --gogottrpc_out=plugins=ttrpc:models/ttrpcmodels models/ttrpcmodels/*.proto

clean:
	go clean

.PHONY: generate generate-proto
