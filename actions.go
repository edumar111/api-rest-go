package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//====MONGO=========================
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		log.Fatal(err)
	}

	return session
}

var collection = getSession().DB("netflix").C("movies")

//===HEAD METHODS======================
func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

//==METHOD ROUTES======================
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from server web Go")
}
func MovieList(w http.ResponseWriter, r *http.Request) {
	log.Println("MovieList")
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}

	responseMovies(w, 200, results)
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)
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
	//log.Println(movie_data)
	err = collection.Insert(movie_data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movie_data)

}
