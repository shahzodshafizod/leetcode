package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestIdealArrays$
func TestIdealArrays(t *testing.T) {
	for _, tc := range []struct {
		n        int
		maxValue int
		count    int
	}{
		{n: 2, maxValue: 5, count: 10},
		{n: 5, maxValue: 3, count: 11},
	} {
		count := idealArrays(tc.n, tc.maxValue)
		assert.Equal(t, tc.count, count)
	}
}
