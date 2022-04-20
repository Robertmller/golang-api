package main

import (
	"context"
	pb "golang-api/grpc/pb"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

type MovieManagementServer struct {
	pb.UnimplementedMovieManagerServer
}

func (s *MovieManagementServer) CreateNewMovie(ctx context.Context, in *pb.NewMovie) (*pb.Movie, error) {
	log.Printf("Received: %v", in.GetTitle())
	var movie_id int32 = int32(rand.Intn(1000))

	return &pb.Movie{Title: in.GetTitle(), ImageUrl: in.GetImageUrl(), Imdb: in.GetImdb(), ReleaseYear: in.GetReleaseYear(), Gender: in.GetGender(), Duration: in.GetDuration(), Director: in.GetDirector(), Id: movie_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faild to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMovieManagerServer(s, &MovieManagementServer{})
	log.Printf("Server listining at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve %v", err)
	}

}
