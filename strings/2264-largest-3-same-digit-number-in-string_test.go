package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLargestGoodInteger$
func TestLargestGoodInteger(t *testing.T) {
	for _, tc := range []struct {
		num  string
		good string
	}{
		{num: "6777133339", good: "777"},
		{num: "2300019", good: "000"},
		{num: "42352338", good: ""},
	} {
		good := largestGoodInteger(tc.num)
		assert.Equal(t, tc.good, good)
	}
}
