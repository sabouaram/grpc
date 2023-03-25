package main

import (
	"context"
	pb "github.com/sabouaram/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:50051"

func doGreet(c pb.GreetServiceClient) {
	log.Printf("goGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Salim",
	})
	if err != nil {
		log.Fatalf("could not greet")
	}
	log.Printf(res.Result)

}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)

}
