generate: generate-proto

generate-proto:
	protoc -I models --go_out=plugins=grpc:models models/*.proto

.PHONY: generate generate-proto
