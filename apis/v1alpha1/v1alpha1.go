//go:generate protoc -I../../protos/v1alpha1 --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative baaa.proto

package v1alpha1
