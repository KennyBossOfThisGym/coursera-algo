package main

import "fmt"

func main() {
	var m int
	coins := []int{10, 5, 1}
	fmt.Scan(&m)
	count := moneyChange(m, coins)
	fmt.Println(count)
}

func moneyChange(m int, coins []int) int {
	var change []int
	for _, coin := range coins {
		for m >= coin {
			change = append(change, coin)
			m = m - coin
		}
	}
	return (len(change))
}
