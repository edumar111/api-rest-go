package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies = Movies{
	Movie{"Sin Limites", 2013, "Desconocido"},
	Movie{"Batman Begin", 2013, "Leo saple"},
	Movie{"Top Gouns", 2013, "George Lucas"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server web Go")
}
func MovieList(w http.ResponseWriter, r *http.Request) {
	log.Println("MovieList")
	//fmt.Fprintf(w, "List Movies")
	json.NewEncoder(w).Encode(movies)
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Show Movie %s", movie_id)
}
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	log.Println("MovieAdd")
	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	log.Println(movie_data)
	movies = append(movies, movie_data)

}
