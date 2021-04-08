package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(FibonacciLast(n))
}

func FibonacciLast(n int) int {
	fibo := []int{0, 1}
	if n < 2 {
		return fibo[n]
	}
	for i := 2; i <= n; i++ {
		fibo = append(fibo, (fibo[i-1]+fibo[i-2])%10)
	}
	return fibo[n]
}
