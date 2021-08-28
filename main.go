package main

import (
    "fmt"
    //"log"
    "net/http"
	"strings"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gocolly/colly"
)

type animeInfo struct {
	Name  string `json:"Name"`
	Rank     string `json:"Rank"`
	Popularity string `json:"Popularity"`
	Members    string`json:"Members"`
	//Score   string
}



func ScrapeWebsite  (id string) animeInfo {

	info := animeInfo{}
	c := colly.NewCollector(
		//colly.AllowedDomains("myanimelist"),
	)
	URL := "https://myanimelist.net/anime/" + id
	c.OnHTML(".title-name",func(e *colly.HTMLElement) {
		    info.Name = e.Text
			fmt.Println(e.Text)
	})

	c.OnHTML(".stats-block", func(e *colly.HTMLElement){

		var StatArr[] string
		e.ForEach("span.numbers", func (_ int ,el *colly.HTMLElement){
			StatArr = append(StatArr, el.Text)
		})

		fmt.Println(StatArr)
		info.Rank = strings.Split(StatArr[0]," ")[1][1:]
		info.Popularity = strings.Split(StatArr[1]," ")[1][1:]
		info.Members = strings.Split(StatArr[2]," ")[1]
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(URL)
	return info
}

func AnimeHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, ok := vars["id"]
	fmt.Println(ok)
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	
	// Web scraping
	info := ScrapeWebsite(id)
	b, err := json.Marshal(info)
	if err != nil {
        fmt.Println(err)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, string(b))
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