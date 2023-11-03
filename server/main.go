package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/qh-org/samples/apis/v1alpha1"
)

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetTitle(ctx context.Context, in *pb.GetTitleReq) (*pb.GetTitleResp, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetTitleResp{
		Baaa: GetTitle(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterInventoryServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GetTitle() *pb.BAAA {
	return &pb.BAAA{
		Title: "title",
	}
}
