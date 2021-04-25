package main

import (
	"log"
	"net"

	"github.com/audrenbdb/algorithm-microsrv"
	"github.com/audrenbdb/algorithm-microsrv/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	srvc := algorithm.NewService()
	fizzBuzzEndpoint := algorithm.CreateEndpoints(srvc)
	grpcServer := algorithm.NewGRPCServer(fizzBuzzEndpoint)
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("unable to start grpc server")
	}

	baseServer := grpc.NewServer()
	reflection.Register(baseServer)
	pb.RegisterSolveServer(baseServer, grpcServer)
	log.Println("Server started")
	baseServer.Serve(grpcListener)
}
