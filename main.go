package main

import (
    "fmt"
    //"log"
    "net/http"
	"strings"
	_"errors"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gocolly/colly"
)

// Struct to store details about the Anime
type animeInfo struct {
	Name  string `json:"Name"`
	Rank     string `json:"Rank"`
	Popularity string `json:"Popularity"`
	Members    string`json:"Members"`
}


// This function is called if the given Anime ID is not found in the database
func ScrapeWebsite  (id string) animeInfo {

	info := animeInfo{}
	c := colly.NewCollector(
		// Allow the crawler to only navigate to links pertaining to the site mentioned
		colly.AllowedDomains("myanimelist.net"),
		colly.CacheDir("./cache"),
	)
	URL := "https://myanimelist.net/anime/" + id


	c.OnHTML(".title-name",func(e *colly.HTMLElement) {
		    info.Name = e.Text
			// Get Anime name
			fmt.Println(e.Text)
	})

	c.OnHTML(".stats-block", func(e *colly.HTMLElement){

		// The rank, popularity and Members of the anime are stored in spans
		// With class name of numbers

		var StatArr[] string
		e.ForEach("span.numbers", func (_ int ,el *colly.HTMLElement){
			StatArr = append(StatArr, el.Text)
		})

		fmt.Println(StatArr)
		info.Rank = strings.Split(StatArr[0]," ")[1][1:]
		info.Popularity = strings.Split(StatArr[1]," ")[1][1:]
		info.Members = strings.Split(StatArr[2]," ")[1]
	})

	c.OnError(func(r *colly.Response, err error) {
		// If anime ID is unavailble on the site
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
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
	
	// Calling the web scraping function since id has not been found
	// in the database
 	info := ScrapeWebsite(id)
	if len(info.Name) == 0 {
		// Send an error to the user
		w.Header().Set("Content-Type", "application/json")
		ErrMsg := `{ "error : Incorrect ID or anime not available on https://myanimelist.net/ }`
		JSONMsg, err := json.Marshal(ErrMsg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, string(JSONMsg))
		return
	}
	
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
	// To handle routes with proper path
	r.HandleFunc("/anime/{id}", AnimeHandler).Methods("GET")
	// TO inform the user in case ID was not included in the path
	r.HandleFunc("/anime", AnimePathErrorHandler).Methods("GET")

	http.ListenAndServe(":3000", r)
}