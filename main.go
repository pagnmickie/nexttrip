package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

//const APIURL = "http://svc.metrotransit.org/NexTrip/"
//const Headers = {"Content-Type": "application/json", "Accept": "application/json"}

// http://svc.metrotransit.org/NexTrip/{ROUTE}/{DIRECTION}/{STOP}/?format=json

func main() {
	//busRoute := GetBusRoute()
	//direction := GetDirection()
	//busStop := BusStop()
	//fmt.Println(busRoute, direction, busStop)
	//url := APIURL + busRoute +"/"+ direction + "/" + busStop + "?format=json"

	resp, err := http.Get("http://svc.metrotransit.org/NexTrip/901/1/TF22?format=json")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle err
		}
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}}

// user enters the name of the route they want to use
//func GetBusRoute() string {
//	println("Enter bus route: ie: 901")
//
//	var input string
//	fmt.Scanln(&input)
//
//	return input
//}
//
//// after the route is selected, the GetDirections() function chooses which direction of travel
//func GetDirection() string {
//	println("Enter route direction: ie: 1")
//
//	var input string
//	fmt.Scanln(&input)
//
//	return input
//}
//
//// after the route and direction are selected, the BusStop() operation finds the trip starting point
//func BusStop() string {
//	println("Enter your bus stop: ie: TF22")
//
//	var input string
//	fmt.Scanln(&input)
//
//	return input
//}


