package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumTilings$
func TestNumTilings(t *testing.T) {
	for _, tc := range []struct {
		n     int
		count int
	}{
		{n: 3, count: 5},
		{n: 1, count: 1},
	} {
		count := numTilings(tc.n)
		assert.Equal(t, tc.count, count)
	}
}
