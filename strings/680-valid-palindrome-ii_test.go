package strings

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestValidPalindrome$
func TestValidPalindrome(t *testing.T) {
	for _, data := range []struct {
		s     string
		valid bool
	}{
		{s: "eedede", valid: true},
		{s: "axbcbaba", valid: false},
		{s: "race a car", valid: true},
		{s: "abccdba", valid: true},
		{s: "abcdefdba", valid: false},
		{s: "", valid: true},
		{s: "a", valid: true},
		{s: "ab", valid: true},
		{s: "abcdcdba", valid: true},
		{s: "abcdfcdba", valid: false},
		{s: "abcdfefcdba", valid: false},
		{s: ".,", valid: true},
		{s: "aba", valid: true},
		{s: "abca", valid: true},
		{s: "abc", valid: false},
		{s: "ebcbbececabbacecbbcbe", valid: true},
		{s: "deddde", valid: true},
		{s: "acucucubucucucua", valid: true},
		{s: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", valid: true},
	} {
		valid := validPalindrome(data.s)
		if !assert.Equal(t, data.valid, valid) {
			log.Printf("ERROR s [%s]\n", data.s)
		}
	}
}
