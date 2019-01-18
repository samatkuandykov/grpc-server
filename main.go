package main

import (
	"fmt"
	"log"
	"net"
	
	"google.golang.org/grpc"
	context "golang.org/x/net/context"
	ns "./names"
	db "./db"
)

const grpcPort = ":8082"
const httpPort = ":8081"

func main() {
	// start listening for grpc
	fmt.Println("Starting listening for gRPC")
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	// Create new grpc server
	server := grpc.NewServer()
	// Register service
	ns.RegisterNameServiceServer(server, new(NameService))
	// Start serving requests
	server.Serve(listen)
}

type NameService struct {}

func (c *NameService) GetAllNames(ctx context.Context, in *ns.Empty) (*ns.Names, error) {
	names := db.GetAllNames()
	return &names, nil
}

func (c *NameService) AddName(ctx context.Context, in *ns.Name) (*ns.Name, error) {
	name := db.AddName(in)
	return name, nil
}