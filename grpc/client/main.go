package main

import (
	"context"
	pb "golang-api/grpc/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:5000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewMovieManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var new_movie = make(map[string]string)

	new_movie["Star Wars"] = "http://test"
	new_movie["Star Wars 2"] = "http://test-2"

	for title, imageUrl := range new_movie {
		r, err := c.CreateNewMovie(ctx, &pb.NewMovie{Title: title, ImageUrl: imageUrl})
		if err != nil {
			log.Fatalf("could not create movie %v", err)
		}
		log.Printf(`movie details
		
		TITLE: %s
		IMAGEURL: %d
		ID: %d`, r.GetTitle(), r.GetImageUrl(), r.GetId())
	}
}
