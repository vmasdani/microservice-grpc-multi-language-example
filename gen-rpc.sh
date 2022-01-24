protoc --go_out=./go --go_opt=paths=source_relative \
    --go-grpc_out=./go --go-grpc_opt=paths=source_relative \
    --proto_path=protos protos/mastermodel.proto &&\

cp -r protos/* spring/src/main/proto
