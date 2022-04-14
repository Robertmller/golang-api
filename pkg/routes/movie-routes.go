package routes

import (
	"golang-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", controllers.DeleteMovie).Methods("DELETE")
}
