package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestLongestIdealString$
func TestLongestIdealString(t *testing.T) {
	for _, tc := range []struct {
		s      string
		k      int
		length int
	}{
		{s: "acfgbd", k: 2, length: 4},
		{s: "abcd", k: 3, length: 4},
		{s: "eduktdb", k: 15, length: 5},
		{s: "a", k: 1, length: 1},
	} {
		length := longestIdealString(tc.s, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
