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

func NewMovieManagementServer() *MovieManagementServer {
	return &MovieManagementServer{
		movie_list: &pb.MovieList{},
	}
}

type MovieManagementServer struct {
	pb.UnimplementedMovieManagerServer
	movie_list *pb.MovieList
}

func (server *MovieManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Faild to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMovieManagerServer(s, server)
	log.Printf("Server listining at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *MovieManagementServer) CreateNewMovie(ctx context.Context, in *pb.NewMovie) (*pb.Movie, error) {
	log.Printf("Received: %v", in.GetTitle())
	var movie_id int32 = int32(rand.Intn(1000))
	created_movie := &pb.Movie{Title: in.GetTitle(), ImageUrl: in.GetImageUrl(), Imdb: in.GetImdb(), ReleaseYear: in.GetReleaseYear(), Gender: in.GetGender(), Duration: in.GetDuration(), Director: in.GetDirector(), Id: movie_id}
	s.movie_list.Movies = append(s.movie_list.Movies, created_movie)
	return created_movie, nil
}

func (s *MovieManagementServer) GetMovies(ctx context.Context, in *pb.GetMoviesParams) (*pb.MovieList, error) {
	return s.movie_list, nil
}

func main() {
	var movie_management_server *MovieManagementServer = NewMovieManagementServer()
	if err := movie_management_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
