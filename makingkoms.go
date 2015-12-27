// 
package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/strava/go.strava"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"strconv"
	"github.com/kristhompson/datastore"
	
)



type Segment struct {
    Name  string    `json:"name"`
    Kom 	bool      `json:"kom"`   
}

type Segments []Segment
var segments Segments



func main() {
	
	datastore.InitDB()
	
	// establish my routes
	r := mux.NewRouter()
	
	r.HandleFunc("/athlete", loadAthlete).Methods("GET")
	r.HandleFunc("/athlete/{athleteId}", loadAthlete).Methods("GET")
	r.HandleFunc("/athleteactivities", loadAtleteActivities).Methods("GET")
	r.HandleFunc("/activityDetails/{activityId}", loadActivityDetails).Methods("GET")
	r.HandleFunc("/segmentLeaderboard/{segmentId}", loadSegmentLeaderboard).Methods("GET")
	
	
	// holding on for example sake
	r.HandleFunc("/goodbye", goodbye)
  
	http.Handle("/", r)
	
	
	handler := cors.Default().Handler(r)
    http.ListenAndServe(":9000", handler)
}


/*
optionall passes in athlete id
*/
func loadAthlete(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8") 
	
	/**
	Get athlete if passed in
	*/
	vars := mux.Vars(r)
    athleteIdStr := vars["athleteId"]
	athleteId, _ := strconv.ParseInt(athleteIdStr, 0, 64)
	fmt.Printf("id is")
	fmt.Printf(athleteIdStr)
	
	accessToken, _ := getStravaConfig()
	client := strava.NewClient(accessToken)
	

	// returns a AthleteDetailed object, the second variable I think is errors
	if(athleteId != 0){
		service := strava.NewAthletesService(client)
		athlete, _ := service.Get(athleteId).Do()
		
		if err := json.NewEncoder(w).Encode(athlete); err != nil {
        panic(err)
    }
	}else{
		service := strava.NewCurrentAthleteService(client)
		athlete, _ := service.Get().Do()
		
		if err := json.NewEncoder(w).Encode(athlete); err != nil {
        panic(err)
    }
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
	
	
	
func loadActivityDetails(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8") 
	
	 vars := mux.Vars(r)
    activityIdStr := vars["activityId"]
	activityId, _ := strconv.ParseInt(activityIdStr, 0, 64)
	fmt.Printf(activityIdStr)
	
	accessToken, _ := getStravaConfig()
	client := strava.NewClient(accessToken)
	service := strava.NewActivitiesService(client)
	
	
	// returns a slice of ActivityDetail objects
	activity, _ := service.Get(activityId).IncludeAllEfforts().Do()
		
	if err := json.NewEncoder(w).Encode(activity); err != nil {
        panic(err)
    }
}
	
	
/**
	Get the leaderboard.  Optional pass in the parameter of 'following'
	
	/segmentLeaderboard/{segmentId?following=true}
	*/	
func loadSegmentLeaderboard(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8") 
	
	 vars := mux.Vars(r)
    segmentIdStr := vars["segmentId"]
	
	
	// get query param if it exists
	followingStr := r.URL.Query().Get("following")
	following, _ := strconv.ParseBool(followingStr)
	
	segmentId, _ := strconv.ParseInt(segmentIdStr, 0, 64)
	
	
	accessToken, _ := getStravaConfig()
	client := strava.NewClient(accessToken)
	service := strava.NewSegmentsService(client)
	
	leaderboard, _ := service.GetLeaderboard(segmentId).Following().Do()
	
	// returns a slice of ActivityDetail objects
	if(following){
		leaderboard, _ = service.GetLeaderboard(segmentId).Following().Do()
	}else{
		leaderboard, _ = service.GetLeaderboard(segmentId).Do()
	}
	
		
	if err := json.NewEncoder(w).Encode(leaderboard); err != nil {
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



	
	

	/**
	Holder of basic information
	*/	
func getStravaConfig()(string, string){
	accessToken := "a94d7f430c0da41b2062ac49ed7ff7e838fc6ec4"
	athleteId := "52931"
	return accessToken, athleteId
}

		

