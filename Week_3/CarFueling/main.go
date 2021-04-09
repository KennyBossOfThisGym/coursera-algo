package main

import "fmt"

func main() {
	var dist, rng, stationNum int
	stations := []int{0} // init stations with starting point 0
	fmt.Scan(&dist)
	fmt.Scan(&rng)
	fmt.Scan(&stationNum)
	for i := 0; i < stationNum; i++ { // O(n)
		var station int
		fmt.Scan(&station)
		stations = append(stations, station)
	}
	stations = append(stations, rng) // add dest point
}

func travel(dist int, rng int, stations []int) int {
	// min_stops := -1
	var min_stops int
	for i := 0; i < len(stations); i++ { //O(n)
		if  
		for i := +1 ; i < len(stations); i++{
	
		}
		min_stops++
	}
	return min_stops
}
