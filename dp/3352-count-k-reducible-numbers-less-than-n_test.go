package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountKReducibleNumbers$
func TestCountKReducibleNumbers(t *testing.T) {
	for _, tc := range []struct {
		s     string
		k     int
		count int
	}{
		{s: "111", k: 1, count: 3},
		{s: "1000", k: 2, count: 6},
		{s: "1", k: 3, count: 0},
		{s: "101", k: 2, count: 4},
	} {
		count := countKReducibleNumbers(tc.s, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
