package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMaxFreqSum$
func TestMaxFreqSum(t *testing.T) {
	for _, tc := range []struct {
		s     string
		total int
	}{
		{s: "successes", total: 6},
		{s: "aeiaeia", total: 3},
	} {
		total := maxFreqSum(tc.s)
		assert.Equal(t, tc.total, total)
	}
}
