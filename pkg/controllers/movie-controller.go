package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang-api/pkg/models"
	"golang-api/pkg/utils"

	"github.com/gorilla/mux"
)

var NewMovie models.Movie

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	newMovies := models.GetAllMovies()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, _ := models.GetMovieById(ID)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParseBody(r, CreateMovie)
	b := CreateMovie.CreateMovie()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movie := models.DeleteMovie(ID)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movie{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, _ := models.GetMovieById(ID)
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateMovie.Author != "" {
		movieDetails.Author = updateMovie.Author
	}
	if updateMovie.Publication != "" {
		movieDetails.Publication = updateMovie.Publication
	}

	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkgplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
