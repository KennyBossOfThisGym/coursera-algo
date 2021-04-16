package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	n := int((-3 + math.Sqrt(float64(9+8*N))) / 2)

	var output []int

	if N > 2 {
		for i := 1; i <= N; i++ {
			output = append(output, i)
			if i == n {
				break
			}
		}
		r := N - n
		output = append(output, r)
	} else {
		output = append(output, N)
	}
	fmt.Println(len(output))
	for _, v := range output {
		fmt.Printf("%v ", v)
	}

}
