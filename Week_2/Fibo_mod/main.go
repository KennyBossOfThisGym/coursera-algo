package main

import "fmt"

func main() {
	var n, mod int
	fmt.Scan(&n)
	fmt.Scan(&mod)
	fmt.Println(FibonacciMod(n, mod))
}

func FibonacciMod(n, mod int) int {
	fibo := []int{0, 1}

	if n < 2 {
		return fibo[n]
	}

	var period int
	var remainder int

	for i := 2; i <= n; i++ {
		f := (fibo[i-1] + fibo[i-2]) % mod // compute current fibo
		f_next := (f + fibo[i-1]) % mod    //compute next fibo

		if f == 0 && f_next == 1 { //check if new period starts -> cut and calculate needed fibo
			period = i
			remainder = n % period
			return fibo[remainder]
		}

		fibo = append(fibo, f) // else append fibo slice
	}
	return fibo[n]
}
