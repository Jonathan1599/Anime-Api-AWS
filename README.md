# Trademarkia Go Anime API (18BIS0100)


## Installation
- Download Go 1.7 from the given [link](https://golang.org/doc/install)
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

## Deployment
This code has been deployed on an AWS EC2 instance which can be accessed on http://3.143.230.153/anime/{id} 
or http://ec2-3-143-230-153.us-east-2.compute.amazonaws.com/anime/{id}
EXAMPLE :  http://3.143.230.153/anime/223 
Note that this instance will be kept online for a maximum duration of **three days**.

## Video code Walk through
 [Link](https://drive.google.com/file/d/1OvJ2NWR_aQ0QWeRqdVNI0lvE4IPlEb_x/view?usp=sharing) <br />
 PS : If you watched the code walk through, I was unable to use CloudFront to cache and act as a CDN to allow for time based caching :grimacing:
