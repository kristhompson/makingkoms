// load-segments-activity.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/strava/go.strava"
	"encoding/json"
	"github.com/gorilla/mux"
	
)

type Segment struct {
    Name      string    `json:"name"`
    Kom bool      `json:"kom"`
    
}

type Segments []Segment
var segments Segments



func main() {
	
	
	// establish my routes
	
	r := mux.NewRouter()
    r.HandleFunc("/t", hello)
	r.Methods("GET")
	
	
	
	http.Handle("/", r)
	
	http.HandleFunc("/goodbye", goodbye)
	http.ListenAndServe(":8000", nil)

  
}

func goodbye(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8") 
	
	segments := Segments{
        {Name: "United States", Kom: true},
        {Name: "Bahamas", Kom: true},
        {Name: "Japan", Kom: false},
    }
	
	 w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(segments); err != nil {
        panic(err)
    }
	
	io.WriteString(w, "Goodbye world!")
}


func helloPost(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world POST!")
	}
	
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	
	var accessToken string
	var segmentId int64
	
	accessToken = "a94d7f430c0da41b2062ac49ed7ff7e838fc6ec4"
	segmentId = 9773190
	
	client := strava.NewClient(accessToken)
	
	//io.WriteString(w,  segmentId.)
	
	segment, err := strava.NewSegmentsService(client).Get(segmentId).Do()
	if err != nil {
		os.Exit(1)
	}
	
	
	
	fmt.Printf(segment.Name)
	
	io.WriteString(w, segment.Name)
	
}
