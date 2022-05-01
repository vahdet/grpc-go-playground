protoc --proto_path=api/proto/v1 \
    --go_out=pkg/api --go_opt=paths=source_relative \
    --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
    person.proto
