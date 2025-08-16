package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaximum69Number$
func TestMaximum69Number(t *testing.T) {
	for _, tc := range []struct {
		num       int
		maximized int
	}{
		{num: 9669, maximized: 9969},
		{num: 9996, maximized: 9999},
		{num: 9999, maximized: 9999},
	} {
		maximized := maximum69Number(tc.num)
		assert.Equal(t, tc.maximized, maximized)
	}
}
