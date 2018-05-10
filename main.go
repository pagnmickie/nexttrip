package main

import (
	"fmt"
	"os"
	"log"
	"./nextrip"
)

const APIURL = "http://svc.metrotransit.org/NexTrip/"

//const Headers = {"Content-Type": "application/json", "Accept": "application/json"}

// http://svc.metrotransit.org/NexTrip/{ROUTE}/{DIRECTION}/{STOP}?format=json
//curl -kv "http://svc.metrotransit.org/NexTrip/94/2/6SHE?format=json"

type NexTripDeparture struct {
	Actual           bool		`json:"Actual,omitempty"`
	BlockNumber      int		`json:"BlockNumber,omitempty"`
	DepartureText    string		`json:"DepartureText,omitempty"`
	DepartureTime    string		`json:"DepartureTime,omitempty"`
	Description      string		`json:"Description,omitempty"`
	Gate             string		`json:"Gate,omitempty"`
	Route            string		`json:"Route,omitempty"`
	RouteDirection   string		`json:"RouteDirection,omitempty"`
	Terminal         string		`json:"Terminal,omitempty"`
	VehicleHeading   int		`json:"VehicleHeading,omitempty"`
	VehicleLatitude  int		`json:"VehicleLatitude,omitempty"`
	VehicleLongitude int		`json:"VehicleLongitude,omitempty"`
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("Please enter three arguments")
	}

	route, err := nextrip.FindRouteByDescription(args[0])
	if err != nil {
		log.Fatal(err)
	}

	direction, err1 := nextrip.FindRouteDirectionByText(route.Route, args[2])
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(direction.Value)
	//busRoute := GetBusRoute()
	//direction := GetDirection()
	//busStop := BusStop()
	//fmt.Println(busRoute, direction, busStop)
	//url := APIURL + busRoute + "/" + direction + "/" + busStop + "?format=json"
	////url := APIURL + busRoute +"/"+ direction + "/" + busStop + "?format=json"
	//
	//resp, err := http.Get(url)
	//if err != nil {
	//	// handle err
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode == http.StatusOK {
	//	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		// handle err
	//	}
	//	//bodyString := string(bodyBytes)
	//	//deptTime := bodyString
	//
	//	r := []NexTripDeparture{}
	//	err = json.Unmarshal(bodyBytes, &r)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(r[0].DepartureTime)
	//}
}

// user enters the name of the route they want to use
func GetBusRoute() string {
	println("Enter bus route: ie: 94 (94 Express)")

	var input string
	fmt.Scanln(&input)

	return input
}

//// after the route is selected, the GetDirections() function chooses which direction of travel
func GetDirection() string {
	println("Enter route direction: ie: 2 (Eastbound)")

	var input string
	fmt.Scanln(&input)

	return input
}

//// after the route and direction are selected, the BusStop() operation finds the trip starting point
func BusStop() string {
	println("Enter your bus stop: ie: 6SHE (6th and Hennepin)")

	var input string
	fmt.Scanln(&input)

	return input
}
