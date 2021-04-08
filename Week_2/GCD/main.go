package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(GCD(a, b))
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	var remainder int
	remainder = a % b
	return GCD(b, remainder)

}
