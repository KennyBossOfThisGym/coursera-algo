package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	n, seq := Input()
	maxProduct := MaxPairWise(n, seq)
	fmt.Println(maxProduct)
}
func MaxPairWise(n int64, seq []int64) int64 {
	var maxProduct int64
	sort.Slice(seq, func(i, j int) bool {
		return seq[i] > seq[j]
	})

	maxProduct = seq[0] * seq[1]

	// if len(seq) > 1 {	// Naive
	// 	for i := 0; i < len(seq)-1; i++ {
	// 		maxProduct = int64(math.Max(float64(maxProduct), float64(seq[i])*float64(seq[i+1])))
	// 	}
	// } else {
	// 	maxProduct = seq[0]
	// }

	return maxProduct
}
func Input() (int64, []int64) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.ParseInt(scanner.Text(), 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	var seq []int64
	seq_input := scanner.Text()
	for _, v := range strings.Fields(seq_input) {
		j, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			log.Fatal(err)
		}

		seq = append(seq, j)

	}
	// if len(seq) != n {
	// 	log.Fatal("Sequence length does not equal " + strconv.Itoa(n))
	// }
	return n, seq
}
