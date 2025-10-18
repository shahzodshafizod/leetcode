package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberOfStableArrays$
func TestNumberOfStableArrays(t *testing.T) {
	for _, tc := range []struct {
		zero  int
		one   int
		limit int
		cnt   int
	}{
		{zero: 1, one: 1, limit: 2, cnt: 2},
		{zero: 1, one: 2, limit: 1, cnt: 1},
		{zero: 3, one: 3, limit: 2, cnt: 14},
	} {
		cnt := numberOfStableArrays(tc.zero, tc.one, tc.limit)
		assert.Equal(t, tc.cnt, cnt)
	}
}
