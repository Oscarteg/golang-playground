package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal("failed to listen: %v", err)
		panic(err)
	}

	server := grpc.NewServer()

}
