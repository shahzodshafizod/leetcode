package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountRoutes$
func TestCountRoutes(t *testing.T) {
	for _, tc := range []struct {
		locations []int
		start     int
		finish    int
		fuel      int
		output    int
	}{
		{locations: []int{2, 3, 6, 8, 4}, start: 1, finish: 3, fuel: 5, output: 4},
		{locations: []int{4, 3, 1}, start: 1, finish: 0, fuel: 6, output: 5},
		{locations: []int{5, 2, 1}, start: 0, finish: 2, fuel: 3, output: 0},
		{locations: []int{2, 1, 5}, start: 0, finish: 0, fuel: 3, output: 2},
		{locations: []int{1, 2, 3}, start: 0, finish: 2, fuel: 40, output: 615088286},
	} {
		output := countRoutes(tc.locations, tc.start, tc.finish, tc.fuel)
		assert.Equal(t, tc.output, output)
	}
}
