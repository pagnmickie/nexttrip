package nextrip

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"errors"
	"strings"
)

type Route struct {
	Description string `json:"Description,omitempty"`
	ProviderID  string `json:"ProviderID,omitempty"`
	Route       string `json:"Route,omitempty"`
}

const routesUrl = "http://svc.metrotransit.org/NexTrip/Routes?format=json"

// Returns route by description
// May return an error if no routes are found
func FindRouteByDescription(description string) (*Route, error) {
	response := getRoutes()
	routes := convertResponseToRoute(response)

	for _, route := range routes {
		if strings.Contains(strings.ToLower(route.Description), strings.ToLower(description)) ||
			strings.Contains(strings.ToLower(route.Route), strings.ToLower(description)) {
			return &route, nil
		}
	}

	return nil, errors.New("Route not found: " + description)
}

// Makes an API call to NexTrip and returns response body
func getRoutes() []byte {
	response, err := http.Get(routesUrl)

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

// Converts response body to a slice of Routes
func convertResponseToRoute(response []byte) []Route {
	routes := []Route{}
	err := json.Unmarshal(response, &routes)

	if err != nil {
		log.Fatal(err)
	}

	return routes
}
