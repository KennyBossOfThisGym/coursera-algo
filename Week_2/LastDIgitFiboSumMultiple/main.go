package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&m)
	fmt.Scan(&n)
	fmt.Println(FibonacciMod(m, n))
}

func FibonacciMod(m, n int) int {
	fibo := []int{0, 1}

	if n < 2 {
		return fibo[n]
	}

	var period int
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

	for _, v := range fibo { // compute period sum

		p_sum = (p_sum + v) % 10
	}

	rem_n := n % period
	rem_m := m % period
	var rem_sum int

	if rem_n >= rem_m { // rem n > rem m

		for i := rem_m; i <= rem_n; i++ {
			rem_sum = (rem_sum + fibo[i]) % 10

		}

	} else { // rem m > rem n
		var m_sum int
		var n_sum int
		for i := rem_m; i <= period-1; i++ {
			m_sum = (m_sum + fibo[i]) % 10
		}
		for i := 0; i <= rem_n; i++ {
			n_sum = (n_sum + fibo[i]) % 10

		}

		rem_sum = (n_sum + m_sum) % 10

	}

	sum = (((n-m)/period)*p_sum + rem_sum) % 10

	return sum
}
