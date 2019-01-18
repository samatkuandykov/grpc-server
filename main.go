package main

import (
	server "github.com/samatkuandykov/grpc-server/server"
)

const grpcPort, httpPort = ":8082", ":8081"

func main() {
	server.Serve(grpcPort, httpPort)
}