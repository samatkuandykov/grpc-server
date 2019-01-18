package main

import (
	"fmt"
	"context"
	
	"google.golang.org/grpc"
	ns "github.com/samatkuandykov/grpc-server/names"
)

var empty ns.Empty
var grpcServer = ":8082"

func main() {
	conn, err := grpc.Dial("localhost" + grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()
	client := ns.NewNameServiceClient(conn)
	// name := ns.Name{Name: "Samat"}
	// client.AddName(context.Background(), &name)
	names, err := client.GetAllNames(context.Background(), &empty)
	if err != nil {
		fmt.Println("error creating grpc cli:", err)
	}
	for _, name := range names.Name {
		fmt.Println(name)
	}
}