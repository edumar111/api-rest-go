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
	router.HandleFunc("/contac", Contact)
	fmt.Println("server listen port 8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server web Go")
}
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "page contac")
}
