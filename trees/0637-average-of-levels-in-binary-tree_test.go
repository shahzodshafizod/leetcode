package trees

import (
	"math"
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestAverageOfLevels$
func TestAverageOfLevels(t *testing.T) {
	for _, tc := range []struct {
		root   []any
		output []float64
	}{
		{root: []any{3, 9, 20, nil, nil, 15, 7}, output: []float64{3.00000, 14.50000, 11.00000}},
		{root: []any{3, 9, 20, 15, 7}, output: []float64{3.00000, 14.50000, 11.00000}},
		{root: []any{1}, output: []float64{1.00000}},
		{root: []any{1, 2, 3, 4, 5}, output: []float64{1.00000, 2.50000, 4.50000}},
	} {
		tree := pkg.MakeTree2(tc.root...)
		output := averageOfLevels(tree)
		assert.Len(t, output, len(tc.output))

		for i := range tc.output {
			assert.Less(t, math.Abs(tc.output[i]-output[i]), 1e-5)
		}
	}
}
