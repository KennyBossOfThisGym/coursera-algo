package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(FibonacciMod(n))
}

func FibonacciMod(n int) int {
	fibo := []int{0, 1}

	if n < 2 {
		return fibo[n]
	}

	var period int
	var n_sum int
	var p_sum int
	var sum int

	for i := 2; i <= n; i++ {
		f := (fibo[i-1] + fibo[i-2]) % 10 // compute current fibo
		f_next := (f + fibo[i-1]) % 10    //compute next fibo

		if f == 0 && f_next == 1 { //check if new period starts ->
			// period = i
			// reps := n / period
			// remainder = n % period
			break
		}

		fibo = append(fibo, f) // else append fibo slice
	}

	period = len(fibo)

	for i, v := range fibo { // compute period + rem sum

		p_sum = (p_sum + v) % 10

		if i == (n % period) {
			n_sum = p_sum
		}
	}

	sum = ((n/period)*p_sum + n_sum) % 10

	return sum
}
