package main

import (
	"golang-api/grpc/pb"
	"golang-api/grpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faild to listen: %v", err)
	}

	newServer := grpc.NewServer()
	pb.RegisterMovieManagerServer(newServer, services.NewMovieManagementServer())
	reflection.Register(newServer)

	if err := newServer.Serve(lis); err != nil {
		log.Fatalf("could not serve: %v", err)
	}

}
