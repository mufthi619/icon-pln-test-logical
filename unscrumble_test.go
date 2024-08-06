package icon_pln_logical

import (
	"fmt"
	"strings"
	"testing"
)

func TestUnscramble(t *testing.T) {
	reverseString := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	unscramble := func(text string) string {
		words := strings.Fields(text)
		unscrambledWords := make([]string, len(words))
		for i, word := range words {
			unscrambledWords[i] = reverseString(word)
		}
		return strings.Join(unscrambledWords, " ")
	}

	listJumbleText := []string{
		"iadab itsap ulalreb",
		"nalub kusutret gnalali",
		"italem irad irigayaj",
	}

	for _, val := range listJumbleText {
		readableText := unscramble(val)

		fmt.Println("Original Text:", val)
		fmt.Println("Readable Text:", readableText)
	}
}
