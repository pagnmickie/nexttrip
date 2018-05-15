package nextrip

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"errors"
	"strings"
)

type Direction struct {
	Text  string `json:"Text,omitempty"`
	Value string `json:"Value,omitempty"`
}

const directionsUrl = "http://svc.metrotransit.org/NexTrip/Directions/%s?format=json"

// Returns route direction for a given route
// May return an error if no directions are found
func FindRouteDirectionByText(routeId string, directionText string) (*Direction, error) {
	response := getDirections(routeId)
	directions := convertResponseToDirection(response)

	for _, direction := range directions {
		if strings.Contains(strings.ToLower(direction.Text), strings.ToLower(directionText)) {
			return &direction, nil
		}
	}

	return nil, errors.New("Direction not found: " + directionText)
}

// Makes an API call to NexTrip and returns response body
func getDirections(routeId string) []byte {
	customUrl := fmt.Sprintf(directionsUrl, routeId)
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

// Converts response body to a slice of Directions
func convertResponseToDirection(response []byte) []Direction {
	directions := []Direction{}
	err := json.Unmarshal(response, &directions)

	if err != nil {
		log.Fatal(err)
	}

	return directions
}
