package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	fmt.Println("server listen port 8080")
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}
