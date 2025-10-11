package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestWays$
func TestWays(t *testing.T) {
	for _, tc := range []struct {
		pizza []string
		k     int
		cnt   int
	}{
		{pizza: []string{"A..", "AAA", "..."}, k: 3, cnt: 3},
		{pizza: []string{"A..", "AA.", "..."}, k: 3, cnt: 1},
		{pizza: []string{"A..", "A..", "..."}, k: 1, cnt: 1},
		{pizza: []string{"...", "...", "..."}, k: 1, cnt: 0},
	} {
		cnt := ways(tc.pizza, tc.k)
		assert.Equal(t, tc.cnt, cnt)
	}
}
