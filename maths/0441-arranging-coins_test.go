package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestArrangeCoins$
func TestArrangeCoins(t *testing.T) {
	for _, tc := range []struct {
		n      int
		output int
	}{
		{n: 5, output: 2},              // rows: [1] [2,3] (4,5 incomplete)
		{n: 8, output: 3},              // rows: [1] [2,3] [4,5,6] (7,8 incomplete)
		{n: 1, output: 1},              // rows: [1]
		{n: 3, output: 2},              // rows: [1] [2,3]
		{n: 0, output: 0},              // no rows
		{n: 10, output: 4},             // rows: [1] [2,3] [4,5,6] [7,8,9,10]
		{n: 2147483647, output: 65535}, // large n
	} {
		output := arrangeCoins(tc.n)
		assert.Equal(t, tc.output, output)
	}
}
