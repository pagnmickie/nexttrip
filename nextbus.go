package main

import (
	"fmt"
	"github.com/pagnmickie/nextbus/nextrip"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//curl -kv "http://svc.metrotransit.org/NexTrip/94/2/6SHE?format=json"

func main() {
	args := os.Args[1:]

	// Error if incorrect number of arguments are provided
	if len(args) != 3 {
		log.Fatal("Please enter the route, stop and direction for NexTrip time")
	}

	routeArg := args[0]
	stopArg := args[1]
	directionArg := args[2]

	// Get route based on description
	route, err := nextrip.FindRouteByDescription(routeArg)
	if err != nil {
		log.Fatal(err)
	}

	// Get direction based on direction text
	direction, err1 := nextrip.FindRouteDirectionByText(route.Route, directionArg)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Get stop based on stop text
	stop, err2 := nextrip.FindRouteStopByText(route.Route, direction.Value, stopArg)
	if err2 != nil {
		log.Fatal(err2)
	}

	// Calls API with arguments provided
	nextDeparture, err3 := nextrip.GetNextDeparture(route.Route, direction.Value, stop.Value)
	if err3 != nil {
		fmt.Println("There are no more scheduled buses")
	} else if nextDeparture.Actual == true {
		departure := strings.Replace(nextDeparture.DepartureText, "Min", "Minutes", 1)
		fmt.Println(departure)
	} else {
		departure := convertJsDate(nextDeparture.DepartureTime)
		now := time.Now()
		difference := departure.Sub(now).Minutes()
		// Round minutes by adding .5 minutes so it won't always round down
		fmt.Printf("%d Minutes\n", int64(difference+0.5))
	}
}

// https://github.com/Skovy/metro-api-go
// convert a EPOCH date representation to a golang Time
func convertJsDate(jsDate string) *time.Time {
	pattern := regexp.MustCompile(`Date\(([0-9]{13})-[0-9]{4}`)
	if matched := pattern.FindStringSubmatch(jsDate); len(matched) > 0 {
		if unixTimestamp, err := strconv.Atoi(matched[1][:10]); err == nil {
			converted := time.Unix(int64(unixTimestamp), 0)
			return &converted
		}
	}

	return nil
}
