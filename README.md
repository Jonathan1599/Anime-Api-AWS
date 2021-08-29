# Trademarkia Go Anime API


## Installation
- Download Go 1.7 from the given [link] {https://golang.org/doc/install}
- Clone this repo to your desired location.
- run `go run main.go`
- Linux users, use `sudo go run main.go`
- Follow the steps below.


## Usage
After you've got the project running, navigate to `localhost:3000/anime/{id}`.
NOTE : This API sends responses in JSON
Enter the Anime ID and if myanimelist.net has the ID, the API will return the
* Name 
* Rank
* Popularity, and
* Members of the Anime
In case the ID doesn't exist on myanimelist.net, an error message will be displayed.

## Video code Walk through
 [Link] (https://drive.google.com/file/d/1OvJ2NWR_aQ0QWeRqdVNI0lvE4IPlEb_x/view?usp=sharing) 
 PS : Was unable to use CloudFront to cache and act as a CDN to allow for time based caching :grimacing:
