package main

import (
	"os"
	"log"
	"./nextrip"
	"fmt"
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
		fmt.Println(nextDeparture.DepartureText)
	} else {
		fmt.Println("Next bus @", nextDeparture.DepartureText)
	}

	//	TODO 1. Make the 'else' return next time in minutes

}
