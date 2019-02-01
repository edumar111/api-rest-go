package main

import (
	"fmt"
	"log"

	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server web Go")
	})
	fmt.Println("server listen port 8080")
	server := http.ListenAndServe(":8080", nil)
	log.Fatal(server)

}
