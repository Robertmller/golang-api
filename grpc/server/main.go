package main

import (
	"context"
	pb "golang-api/grpc/pb"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	port = ":5000"
)

func NewMovieManagementServer() *MovieManagementServer {
	return &MovieManagementServer{}
}

type MovieManagementServer struct {
	pb.UnimplementedMovieManagerServer
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

	readBytes, err := ioutil.ReadFile("movies.json")

	var movie_list *pb.MovieList = &pb.MovieList{}

	var movie_id int32 = int32(rand.Intn(1000))

	created_movie := &pb.Movie{Title: in.GetTitle(), ImageUrl: in.GetImageUrl(), Imdb: in.GetImdb(), ReleaseYear: in.GetReleaseYear(), Gender: in.GetGender(), Duration: in.GetDuration(), Director: in.GetDirector(), Id: movie_id}

	if err != nil {
		if os.IsNotExist(err) {
			log.Print("file not found. Creating a new file")
			movie_list.Movies = append(movie_list.Movies, created_movie)

			jsonBytes, err := protojson.Marshal(movie_list)
			if err != nil {
				log.Fatalf("Json Marshaling failed: %v", err)
			}
			if err := ioutil.WriteFile("movies.json", jsonBytes, 0664); err != nil {
				log.Fatalf("failed write to file: %v", err)
			}

			return created_movie, nil
		} else {
			log.Fatalf("error reading file: %v", err)
		}
	}
	if err := protojson.Unmarshal(readBytes, movie_list); err != nil {
		log.Fatalf("failed to parse movies list: %v", err)
	}
	movie_list.Movies = append(movie_list.Movies, created_movie)

	jsonBytes, err := protojson.Marshal(movie_list)
	if err != nil {
		log.Fatalf("Json Marshaling failed: %v", err)
	}
	if err := ioutil.WriteFile("movies.json", jsonBytes, 0664); err != nil {
		log.Fatalf("failed write to file: %v", err)
	}
	return created_movie, nil
}

func (s *MovieManagementServer) GetMovies(ctx context.Context, in *pb.GetMoviesParams) (*pb.MovieList, error) {
	jsonBytes, err := ioutil.ReadFile("movies.json")
	if err != nil {
		log.Fatalf("failed reading from file: %v", err)
	}
	var movie_list *pb.MovieList = &pb.MovieList{}
	if err := protojson.Unmarshal(jsonBytes, movie_list); err != nil {
		log.Fatalf("Unmarshaling failed: %v", err)
	}
	return movie_list, nil
}

func main() {
	var movie_management_server *MovieManagementServer = NewMovieManagementServer()
	if err := movie_management_server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
