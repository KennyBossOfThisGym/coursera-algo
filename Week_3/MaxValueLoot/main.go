package main

import (
	"fmt"
	"sort"
)

type product struct {
	value    float64
	weight   float64
	calories float64
}

type backpack struct {
	products []product
	capacity float64
	value    float64
}

func main() {
	var n int
	var products []product
	var sack backpack
	fmt.Scan(&n)
	fmt.Scan(&sack.capacity)
	for i := 0; i < n; i++ {
		var p product
		fmt.Scan(&p.value)
		fmt.Scan(&p.weight)
		p.calories = p.value / p.weight
		products = append(products, p)
	}

	fillBackPack(products, &sack)
	fmt.Println(sack.value)

}

func fillBackPack(products []product, sack *backpack) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].calories > products[j].calories
	})

	for _, v := range products {
		if sack.capacity >= v.weight {
			sack.products = append(sack.products, v)
			sack.capacity = sack.capacity - v.weight
			sack.value = sack.value + v.value
		} else {
			ratio := sack.capacity / v.weight
			sack.products = append(sack.products, product{value: ratio * v.value, weight: ratio * v.weight, calories: ratio * v.calories})
			sack.capacity = sack.capacity - v.weight*ratio
			sack.value = sack.value + v.value*ratio
		}
		if sack.capacity == 0 {
			return
		}
	}
}
