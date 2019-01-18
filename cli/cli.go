package main

import (
	"fmt"
	"context"
	
	"google.golang.org/grpc"
	ns "../names"
)

var empty ns.Empty

func main() {
	grpcServer := "localhost:8082"
	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
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