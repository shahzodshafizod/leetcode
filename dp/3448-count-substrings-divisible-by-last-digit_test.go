package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountSubstringsDivisible$
func TestCountSubstringsDivisible(t *testing.T) {
	for _, tc := range []struct {
		s     string
		count int64
	}{
		{s: "12936", count: 11},
		{s: "5701283", count: 18},
		{s: "1010101010", count: 25},
		{s: "1212", count: 10},
		{s: "123", count: 5},
	} {
		count := countSubstrings3448(tc.s)
		assert.Equal(t, tc.count, count)
	}
}
