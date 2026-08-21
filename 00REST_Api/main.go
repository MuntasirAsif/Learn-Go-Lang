package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	artiles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	fmt.Fprintf(w, "Endpoint hit: All Articles")
	json.NewEncoder(w).Encode(artiles)
}

func homePage(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "Home page endpoint hit")
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	handleRequests()
}