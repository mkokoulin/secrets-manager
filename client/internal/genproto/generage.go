package genproto

//go:generate protoc --go_out=../pb --go-grpc_out=../pb -I../../api ../../api/users.proto ../../api/secrets.proto
