package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type NoFlyZone struct {
	SpatialReference SpatialReference `json:"spatialReference"`
	Rings            [1][8][2]float64 `json:"rings"`
}

type SpatialReference struct {
	Wkid int `json:"wkid"`
}

var noFlyZonesArray []NoFlyZone

func init() {
	const SPATIAL_REFERENCE_WELL_KNOWN_ID = 102100
	var dtwRings = [1][8][2]float64{
		{
			{-9278977.502393615, 5196972.662366206},
			{-9278404.224681476, 5197240.191965203},
			{-9274505.936238931, 5195673.232885358},
			{-9275518.726863708, 5190055.1113064},
			{-9278881.956108259, 5189061.429938688},
			{-9280869.318843672, 5188660.135540191},
			{-9282646.479751302, 5192481.986954449},
			{-9278977.502393615, 5196972.662366206},
		},
	}

	noFlyZonesArray = []NoFlyZone{
		{
			SpatialReference{
				Wkid: SPATIAL_REFERENCE_WELL_KNOWN_ID,
			},
			dtwRings,
		}}
}

func getAllNoFlyZonesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint '/noFlyZones' hit")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(noFlyZonesArray)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/noFlyZones", getAllNoFlyZonesHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), myRouter))
}

func main() {
	handleRequests()
}
