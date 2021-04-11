package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int // input n
	fmt.Scan(&n)

	var profits []int // input profits times n
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		profits = append(profits, a)
	}

	var clicks []int // input clicks times n
	for i := 0; i < n; i++ {
		var c int
		fmt.Scan(&c)
		clicks = append(clicks, c)
	}

	sum := MaxAdvRevenue(profits, clicks)
	fmt.Println(sum)

}

func MaxAdvRevenue(profits, clicks []int) int {
	sort.Slice(profits, func(i, j int) bool { // sort in- decreasing order
		return profits[i] > profits[j]
	})

	sort.Slice(clicks, func(i, j int) bool {
		return clicks[i] > clicks[j]
	})

	var sum int // calculate optimized sum
	for i, v := range profits {
		sum = sum + (v * clicks[i])
	}
	return sum
}
