package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestTribonacci$
func TestTribonacci(t *testing.T) {
	for _, tc := range []struct {
		n  int
		tn int
	}{
		{n: 4, tn: 4},
		{n: 25, tn: 1389537},
		{n: 37, tn: 2082876103},
	} {
		tn := tribonacci(tc.n)
		assert.Equal(t, tc.tn, tn)
	}
}
