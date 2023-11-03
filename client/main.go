package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/qh-org/samples/apis/v1alpha1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)
	title, err := client.GetTitle(context.Background(), &pb.GetTitleReq{})
	if err != nil {
		log.Fatalf("failed to get title: %v", err)
	}
	fmt.Println(title)
}
