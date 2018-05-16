package nextrip

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// stop struct
type Stop struct {
	Text  string `json:"Text,omitempty"`
	Value string `json:"Value,omitempty"`
}

// stop url
const stopsUrl = "http://svc.metrotransit.org/NexTrip/Stops/%s/%s?format=json"

// export FindRouteStopByText (routeId, directionId)
func FindRouteStopByText(routeId string, directionId string, stopText string) (*Stop, error) {
	response := getStops(routeId, directionId)
	stops := convertResponseToStop(response)

	for _, stop := range stops {
		if strings.Contains(strings.ToLower(stop.Text), strings.ToLower(stopText)) ||
			strings.Contains(strings.ToLower(stop.Value), strings.ToLower(stopText)) {
			return &stop, nil
		}
	}

	return nil, errors.New("Stop not found: " + stopText)
}

func getStops(routeId string, directionId string) []byte {
	customUrl := fmt.Sprintf(stopsUrl, routeId, directionId)
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
		log.Fatal("Unable to get stops")
	}

	return nil
}

func convertResponseToStop(response []byte) []Stop {
	stops := []Stop{}
	err := json.Unmarshal(response, &stops)

	if err != nil {
		log.Fatal(err)
	}

	return stops
}
