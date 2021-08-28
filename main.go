package main

import (
    "fmt"
    //"log"
    "net/http"
	"github.com/gorilla/mux"
)

func AnimeHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, ok := vars["id"]
	fmt.Println(ok)
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	
	// Web scraping

    fmt.Fprintf(w, id)
    fmt.Println("Endpoint Hit: AnimeHandler")
}

func AnimePathErrorHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Please enter an id")
	fmt.Println("Endpoint Hit: AnimePathErrorHandler")
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/anime/{id}", AnimeHandler).Methods("GET")
	r.HandleFunc("/anime", AnimePathErrorHandler).Methods("GET")

	http.ListenAndServe(":3000", r)
}