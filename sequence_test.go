package icon_pln_logical

import (
	"fmt"
	"testing"
)

func TestSequence(t *testing.T) {
	generateSequence := func(n int) []int {
		if n <= 0 {
			return nil
		}
		if n == 1 {
			return []int{0}
		}
		if n == 2 {
			return []int{0, 1}
		}

		seq := make([]int, n)
		seq[0] = 0
		seq[1] = 1

		for i := 2; i < n; i++ {
			seq[i] = seq[i-1] + seq[i-2]
		}

		return seq
	}

	//Ex
	n := 10
	sequence := generateSequence(n)
	for _, num := range sequence {
		fmt.Println(num)
	}
}
