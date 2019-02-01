package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/movies", MovieList)
	router.HandleFunc("/movies/{id}", MovieShow)
	fmt.Println("server listen port 8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server web Go")
}
func MovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List Movies")
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]
	fmt.Fprintf(w, "Show Movie %s", movie_id)
}
