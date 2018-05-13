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
	if len(args) != 3 {
		log.Fatal("Please enter three arguments")
	}

	routeArg := args[0]
	stopArg := args[1]
	directionArg := args[2]

	route, err := nextrip.FindRouteByDescription(routeArg)
	if err != nil {
		log.Fatal(err)
	}

	direction, err1 := nextrip.FindRouteDirectionByText(route.Route, directionArg)
	if err1 != nil {
		log.Fatal(err1)
	}

	stop, err2 := nextrip.FindRouteStopByText(route.Route, direction.Value, stopArg)
	if err2 != nil {
		log.Fatal(err2)
	}

	nextTime, _ := nextrip.GetNextDeparture(route.Route, direction.Value, stop.Value)
	fmt.Println(nextTime.DepartureText)

}
