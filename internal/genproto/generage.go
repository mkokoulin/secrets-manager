package genproto

//go:generate protoc --go_out=. --go-grpc_out=. -I../../api ../../api/users.proto ../../api/secrets.proto
