package router

import (
	"github.com/gorilla/mux"
	"github.com/muntasirashif/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	router.HandleFunc("/students", controller.CreateStudent).Methods("POST")
	router.HandleFunc("/students", controller.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", controller.GetStudent).Methods("GET")
	router.HandleFunc("/students/{id}", controller.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", controller.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students", controller.DeleteAllStudents).Methods("DELETE")

	return router
}
