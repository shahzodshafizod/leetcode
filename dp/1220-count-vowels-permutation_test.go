package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountVowelPermutation$
func TestCountVowelPermutation(t *testing.T) {
	for _, tc := range []struct {
		n     int
		count int
	}{
		{n: 1, count: 5},
		{n: 2, count: 10},
		{n: 5, count: 68},
		{n: 144, count: 18208803},
	} {
		count := countVowelPermutation(tc.n)
		assert.Equal(t, tc.count, count)
	}
}
