package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	context "golang.org/x/net/context"
	ns "github.com/samatkuandykov/grpc-server/names"
	storage "github.com/samatkuandykov/grpc-server/storage"
)

func Serve(grpcPort string, httpPort string) {
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
	go server.Serve(listen)
	
	// Start HTTP server for Rest
	run(grpcPort, httpPort)
}

func run(grpcPort string, httpPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := ns.RegisterNameServiceHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts)
	if err != nil {
		return err
	}

	log.Printf("Starting HTTP/Rest gateway on port %s\n", httpPort)
	return http.ListenAndServe(httpPort, mux)
}

type NameService struct {}

func (c *NameService) GetAllNames(ctx context.Context, in *ns.Empty) (*ns.Names, error) {
	names := storage.GetAllNames()
	return &names, nil
}

func (c *NameService) AddName(ctx context.Context, in *ns.Name) (*ns.Name, error) {
	name := storage.AddName(in)
	return name, nil
}