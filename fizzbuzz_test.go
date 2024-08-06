package icon_pln_logical

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fizzBuzzOutput := func() string {
		var buf bytes.Buffer
		for i := 1; i <= 100; i++ {
			output := ""
			if i%3 == 0 {
				output += "Fizz"
			}
			if i%5 == 0 {
				output += "Buzz"
			}
			if output == "" {
				output = fmt.Sprintf("%d", i)
			}
			buf.WriteString(output + "\n")
		}
		return buf.String()
	}

	expectedOutput := ""
	for i := 1; i <= 100; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}
		expectedOutput += output + "\n"
	}

	got := fizzBuzzOutput()
	if got != expectedOutput {
		t.Errorf("fizzBuzzOutput() = %v, want %v", got, expectedOutput)
	}
}
