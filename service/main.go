package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

const(
	DISTANCE = "200km"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// var first int = (some int)
// first := (some int)
// initialization

// true, false, string
// var func
// %T, %v

type Post struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User string `json:"user"`
	Message string `json:"message"`
	Location Location `json:"location"`
}

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
		return
	}
	fmt.Fprintf(w, "Post recieved: %v\n", p.Message)
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}

	fmt.Fprintf(w, "Search received: %f %f %s \n", lat, lon, ran)


	decoder := json.NewDecoder(r.Body)
	var pin Post
	if err := decoder.Decode(&pin); err != nil {
		panic(err)
		return
	}

	var p Post
	p = Post{
		User: pin.User,
		Message: pin.Message,
		Location: pin.Location,
	}

	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//lat := r.URL.Query().Get("lat")
	//lon := r.URL.Query().Get("lon")
	// fmt.Fprintf(w, "Search received: %s %s", lat, lon)
}