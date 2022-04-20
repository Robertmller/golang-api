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

	new_movie["Star-Wars: A new Hope"] = "https://meufilme.com/starwars4"
	new_movie["Star-Wars: The Empire Strike Back"] = "https://meufilme.com/starwars5"
	new_movie["Star-Wars: Jedi Return"] = "https://meufilme.com/starwars6"

	for title, imageUrl := range new_movie {
		r, err := c.CreateNewMovie(ctx, &pb.NewMovie{Title: title, ImageUrl: imageUrl})
		if err != nil {
			log.Fatalf("could not create movie %v", err)
		}
		log.Printf(`movie details
		
		TITLE: %s
		IMAGEURL: %s
		IMDB: %g
		RELEASEYEAR: %d
		GENDER: %s
		DURATION: %s
		DIRECTOR: %s
		ID: %d`, r.GetTitle(), r.GetImageUrl(), r.GetImdb(), r.GetReleaseYear(), r.GetGender(), r.GetDuration(), r.GetDirector(), r.GetId())
	}
}
