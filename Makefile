.PHONY: proto proto-clean

PROTO_PATH := api
OUT_PATH := internal/generated

proto:
	@echo Generating protobuf...
	if not exist $(OUT_PATH) mkdir $(OUT_PATH)

	protoc -I $(PROTO_PATH) -I . --go_out=$(OUT_PATH) --go-grpc_out=$(OUT_PATH) --grpc-gateway_out=$(OUT_PATH) --validate_out=lang=go:$(OUT_PATH) $(PROTO_PATH)/go_load/go_load.proto

proto-clean:
	@echo Cleaning generated files...
	if exist $(OUT_PATH) rmdir /s /q $(OUT_PATH)