package main

import "testing"

func Test_moneyChange(t *testing.T) {
	coins := []int{10, 5, 1}
	var result int
	result = moneyChange(5, coins)
	if result != 1 {
		t.Errorf("result: %v - must be %v", result, 1)
	}
	result = moneyChange(1, coins)
	if result != 1 {
		t.Errorf("result: %v - must be %v", result, 1)
	}
	result = moneyChange(10, coins)
	if result != 1 {
		t.Errorf("result: %v - must be %v", result, 1)
	}
	result = moneyChange(15, coins)
	if result != 2 {
		t.Errorf("result: %v - must be %v", result, 2)
	}
	result = moneyChange(28, coins)
	if result != 6 {
		t.Errorf("result: %v - must be %v", result, 6)
	}

}
