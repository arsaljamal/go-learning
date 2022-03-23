package main

import (
	svccore "github.com/go_learning/core"
	svcgrpc "github.com/go_learning/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// configure our core service
	userService := svccore.NewService()
	// configure our gRPC service controller
	userServiceController := NewUserServiceController(userService)
	// start a gRPC server
	server := grpc.NewServer()
	svcgrpc.RegisterUserServiceServer(server, userServiceController)
	reflection.Register(server)

	con, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		panic(err)
	}

	log.Printf("Starting gRPC user service on %s...\n", con.Addr().String())

	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
