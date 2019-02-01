package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server web Go")
}
func MovieList(w http.ResponseWriter, r *http.Request) {

	movies := Movies{
		Movie{"Sin Limites", 2013, "Desconocido"},
		Movie{"Batman Begin", 2013, "Leo saple"},
		Movie{"Top Gouns", 2013, "George Lucas"},
	}

	//fmt.Fprintf(w, "List Movies")
	json.NewEncoder(w).Encode(movies)
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Show Movie %s", movie_id)
}
