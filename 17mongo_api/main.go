package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/muntasirashif/mongoapi/controller"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students", controller.GetAllStudents).Methods("GET")
	r.HandleFunc("/students", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", controller.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", controller.DeleteStudent).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5000", "http://localhost:8080", "http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
