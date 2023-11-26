generate_grpc_code:
	protoc \
  --go_out=idm \
  --go_opt=paths=source_relative \
  --go-grpc_out=idm \
  --go-grpc_opt=paths=source_relative \
  idm.proto \
  proto_files/user.proto \
  proto_files/helper.proto
