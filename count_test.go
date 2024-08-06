package icon_pln_logical

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCount(t *testing.T) {
	countNumbers := func(input []string) int {
		count := 0
		for _, item := range input {
			if _, err := strconv.Atoi(item); err == nil {
				count++
			}
		}
		return count
	}

	examples := [][]string{
		{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"},
		{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"},
		{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "91"},
	}

	for _, example := range examples {
		fmt.Println("Input:", example, "Output:", countNumbers(example))
	}
}
