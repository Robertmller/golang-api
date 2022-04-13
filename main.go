package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Title    string    `json: "title"`
	Imdb     string    `json: "imdb"`
	Director *Director `json: "Director"`
}

type Director struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

var movies []Movie

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return

		}

	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

			return

		}
	}
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Fight Club", Imdb: "8.0", Director: &Director{FirstName: "David", LastName: "Fincher"}})
	movies = append(movies, Movie{ID: "2", Title: "The God father", Imdb: "8.0", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "3", Title: "Forest Gump", Imdb: "9.0", Director: &Director{FirstName: "Robert", LastName: "Zemeckis"}})
	movies = append(movies, Movie{ID: "4", Title: "Pulp Fiction", Imdb: "8.5", Director: &Director{FirstName: "Quentin", LastName: "Tarantino"}})
	movies = append(movies, Movie{ID: "5", Title: "Interestelar", Imdb: "8.9", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")

	r.HandleFunc("/movie/{id}", getMovieById).Methods("GET")

	r.HandleFunc("/movie", createMovie).Methods("POST")

	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
