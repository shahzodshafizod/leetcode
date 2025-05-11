package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountBalancedPermutations$
func TestCountBalancedPermutations(t *testing.T) {
	for _, tc := range []struct {
		num   string
		count int
	}{
		{num: "123", count: 2},
		{num: "112", count: 1},
		{num: "12345", count: 0},
		{num: "0101", count: 4},
		{num: "325419", count: 72},
		// {num: "96509861244547846", count: 851350267},
		// {num: "9650986124454784696509861244547846", count: 24468941},
		// {num: "07162132741696526558195913376589552512680680143682957878451885242156416378132712", count: 285771532},
		// {num: "01010101010101010101010101010101010101010101010101010101010101010101010101010101", count: 432969423},
		// {num: "00000000000000000000000000000000000000000000000000000000000000000000000000000000", count: 1},
	} {
		count := countBalancedPermutations(tc.num)
		assert.Equal(t, tc.count, count)
	}
}
