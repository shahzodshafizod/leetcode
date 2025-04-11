package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberOfPowerfulInt$
func TestNumberOfPowerfulInt(t *testing.T) {
	for _, tc := range []struct {
		start  int64
		finish int64
		limit  int
		s      string
		count  int64
	}{
		{start: 1, finish: 6000, limit: 4, s: "124", count: 5},
		{start: 15, finish: 215, limit: 6, s: "10", count: 2},
		{start: 1000, finish: 2000, limit: 4, s: "3000", count: 0},
	} {
		count := numberOfPowerfulInt(tc.start, tc.finish, tc.limit, tc.s)
		assert.Equal(t, tc.count, count)
	}
}
