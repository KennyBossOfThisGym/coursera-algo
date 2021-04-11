package main

import (
	"fmt"
	"sort"
)

func main() {

	var dist int
	fmt.Scan(&dist)

	var carRange int
	fmt.Scan(&carRange)

	var n int
	fmt.Scan(&n)

	var stations []int
	for i := 0; i < n; i++ { // O(n)	// add gas stations
		var station int
		fmt.Scan(&station)
		stations = append(stations, station)
	}
	counter := travel(dist, carRange, stations)
	fmt.Println(counter)

}

func travel(dist int, carRange int, stations []int) int {

	// stations = append(stations, dist) // add finish
	sort.Sort(sort.Reverse(sort.IntSlice(stations)))

	currentPoint := 0
	counter := 0
OUTER:
	for carRange < dist-currentPoint { // while not enough fuel  to reach finish
		for _, station := range stations {

			if station <= currentPoint+carRange { // take first max that fits
				if currentPoint == station { // if doesnt point doesnt change > stale exit
					return -1
				}
				currentPoint = station
				counter++
				continue OUTER // break the search for currrent position
			}

		}

	}

	return counter
}
