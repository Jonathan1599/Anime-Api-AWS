package main

import (
    "fmt"
    //"log"
    "net/http"
	"strings"
	"reflect"
	"github.com/gorilla/mux"
	"github.com/gocolly/colly"
)

type animeInfo struct {
	Name  string 
	Rank     string 
	Popularity string 
	Score   string
}



func ScrapeWebsite(id string){

	info := animeInfo{}
	c := colly.NewCollector(
		//colly.AllowedDomains("myanimelist"),
	)
	URL := "https://myanimelist.net/anime/" + id
	c.OnHTML(".title-name",func(e *colly.HTMLElement) {
		    info.Name = e.Text
			fmt.Println(e.Text)
	})
	
	c.OnHTML(".numbers", func(e *colly.HTMLElement){
		ProcessStr := e.Text
		StrArr := strings.Split(ProcessStr,"\n")
		fmt.Println(reflect.TypeOf(StrArr))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(URL)

}

func AnimeHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, ok := vars["id"]
	fmt.Println(ok)
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	
	// Web scraping
	ScrapeWebsite(id)
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