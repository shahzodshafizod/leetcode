package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinDays$
func TestMinDays(t *testing.T) {
	for _, tc := range []struct {
		n    int
		days int
	}{
		{n: 6, days: 3},
		{n: 10, days: 4},
		// {n: 1073741826, days: 32},
		// {n: 2000000000, days: 32},
	} {
		days := minDays(tc.n)
		assert.Equal(t, tc.days, days)
	}
}
