SERVICE_PATH=./service/
PROTO_PATH=./proto/
PROTO_FILE_PATH=./proto/*.proto
build_proto: 
	protoc --proto_path=${PROTO_PATH} --go_out=${SERVICE_PATH} --go-grpc_out=${SERVICE_PATH} ${PROTO_FILE_PATH}