package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int // input n
	fmt.Scan(&n)

	var stints [][]int // input stints times n
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		stints = append(stints, []int{a, b})
	}
	targetStint := CollectSgn(stints)
	fmt.Println(len(targetStint))
	for _, v := range targetStint {
		fmt.Printf("%v ", v[1])
	}

}

func CollectSgn(stints [][]int) [][]int {
	sort.Slice(stints, func(i, j int) bool { // sort right edge in decreasing order
		return stints[i][1] > stints[j][1]
	})

	targetStint := [][]int{stints[len(stints)-1]}
	for i := len(stints) - 1; i >= 0; i-- { // for each stint

		if stints[i][1] >= targetStint[len(targetStint)-1][1] && stints[i][0] <= targetStint[len(targetStint)-1][1] {
			stints = RemoveSliceElement(stints, i)
		} else {
			targetStint = append(targetStint, stints[i])
			continue
		}

	}
	return targetStint
}

func RemoveSliceElement(slice [][]int, s int) [][]int { // for fuck sake add the slice.remove function
	if len(slice) < 2 {
		slice = slice[:0]
	} else {
		slice = append(slice[:s], slice[s+1:]...)
	}
	return slice

}
