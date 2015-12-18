// 
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/strava/go.strava"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	
)

type Segment struct {
    Name  string    `json:"name"`
    Kom 	bool      `json:"kom"`
    
}

type Segments []Segment
var segments Segments





func main() {
	
	
	// establish my routes
	
	r := mux.NewRouter()
	
	r.HandleFunc("/t", hello).Methods("GET")
	r.HandleFunc("/athlete", loadAthlete).Methods("GET")
	r.HandleFunc("/athleteactivities", loadAtleteActivities).Methods("GET")
	
	
	// holding on for example sake
	r.HandleFunc("/goodbye", goodbye)
  
	http.Handle("/", r)
	
	
	handler := cors.Default().Handler(r)
    http.ListenAndServe(":9000", handler)
  
}



func loadAthlete(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8") 
	
	accessToken, _ := getStravaConfig()
	client := strava.NewClient(accessToken)
	service := strava.NewCurrentAthleteService(client)

	// returns a AthleteDetailed object, the second variable I think is errors
	athlete, _ := service.Get().Do()
	
	
	if err := json.NewEncoder(w).Encode(athlete); err != nil {
        panic(err)
    }
	
}
	
	
	
func loadAtleteActivities(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8") 
	
	accessToken, _ := getStravaConfig()
	client := strava.NewClient(accessToken)
	service := strava.NewCurrentAthleteService(client)
	
	
	// returns a slice of ActivitySummary objects
	activities, _ := service.ListActivities().Do()
		
		if err := json.NewEncoder(w).Encode(activities); err != nil {
        panic(err)
    }
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
	
	
	
func getStravaConfig()(string, string){
	accessToken := "a94d7f430c0da41b2062ac49ed7ff7e838fc6ec4"
	athleteId := "52931"
	return accessToken, athleteId
}

		
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	
	accessToken, _ := getStravaConfig()
	var segmentId int64
	
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