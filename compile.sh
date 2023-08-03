
go mod init github.com/lbaldwin123/testgrpcapp

protoc --go-grpc_out=. testgrpcapp.proto

go run server/server.go
go run client/client.go