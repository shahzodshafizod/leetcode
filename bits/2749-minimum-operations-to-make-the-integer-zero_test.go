package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMakeTheIntegerZero$
func TestMakeTheIntegerZero(t *testing.T) {
	for _, tc := range []struct {
		num1 int
		num2 int
		cnt int
	}{
		{num1: 3, num2: -2, cnt: 3},
		{num1: 5, num2: 7, cnt: -1},
	} {
		cnt := makeTheIntegerZero(tc.num1, tc.num2)
		assert.Equal(t, tc.cnt, cnt)
	}
}
