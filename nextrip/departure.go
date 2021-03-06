package nextrip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Departure struct {
	Actual           bool    `json:"Actual,omitempty"`
	BlockNumber      int     `json:"BlockNumber,omitempty"`
	DepartureText    string  `json:"DepartureText,omitempty"`
	DepartureTime    string  `json:"DepartureTime,omitempty"`
	Description      string  `json:"Description,omitempty"`
	Gate             string  `json:"Gate,omitempty"`
	Route            string  `json:"Route,omitempty"`
	RouteDirection   string  `json:"RouteDirection,omitempty"`
	Terminal         string  `json:"Terminal,omitempty"`
	VehicleHeading   int     `json:"VehicleHeading,omitempty"`
	VehicleLatitude  float64 `json:"VehicleLatitude,omitempty"`
	VehicleLongitude float64 `json:"VehicleLongitude,omitempty"`
}

const departuresUrl = "http://svc.metrotransit.org/NexTrip/%s/%s/%s?format=json"

// Returns next departure time for a given route, direction and stop
// May return an error if no departures are found
func GetNextDeparture(routeId string, directionId string, stopId string) (*Departure, error) {
	response := getDepartures(routeId, directionId, stopId)
	departures := convertResponseToDepartures(response)

	if departures != nil && len(departures) > 0 {
		return &departures[0], nil
	}

	return nil, errors.New("departures not found")
}

// Makes an API call to NexTrip and returns response body
func getDepartures(routeId string, directionId string, stopId string) []byte {
	customUrl := fmt.Sprintf(departuresUrl, routeId, directionId, stopId)
	response, err := http.Get(customUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		return bytes
	} else {
		log.Fatal("Unable to get routes")
	}

	return nil
}

// Converts response body to a slice of Departures
func convertResponseToDepartures(response []byte) []Departure {
	departures := []Departure{}
	err := json.Unmarshal(response, &departures)

	if err != nil {
		log.Fatal(err)
	}

	return departures
}
