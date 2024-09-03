package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestGetLucky$
func TestGetLucky(t *testing.T) {
	for _, tc := range []struct {
		s     string
		k     int
		lucky int
	}{
		{s: "iiii", k: 1, lucky: 36},
		{s: "leetcode", k: 2, lucky: 6},
		{s: "zbax", k: 2, lucky: 8},
		{s: "vbyytoijnbgtyrjlsc", k: 2, lucky: 15},
		{s: "hvmhoasabaymnmsd", k: 1, lucky: 79},
	} {
		lucky := getLucky(tc.s, tc.k)
		assert.Equal(t, tc.lucky, lucky)
	}
}
