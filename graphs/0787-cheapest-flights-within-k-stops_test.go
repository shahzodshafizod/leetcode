package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindCheapestPrice$
func TestFindCheapestPrice(t *testing.T) {
	for _, tc := range []struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		price   int
	}{
		{
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
			src:     0, dst: 3, k: 1, price: 700,
		},
		{
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0, dst: 2, k: 1, price: 200,
		},
		{
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0, dst: 2, k: 0, price: 500,
		},
		{
			n:       11,
			flights: [][]int{{0, 3, 3}, {3, 4, 3}, {4, 1, 3}, {0, 5, 1}, {5, 1, 100}, {0, 6, 2}, {6, 1, 100}, {0, 7, 1}, {7, 8, 1}, {8, 9, 1}, {9, 1, 1}, {1, 10, 1}, {10, 2, 1}, {1, 2, 100}},
			src:     0, dst: 2, k: 4, price: 11,
		},
		{
			n:       10,
			flights: [][]int{{3, 4, 4}, {2, 5, 6}, {4, 7, 10}, {9, 6, 5}, {7, 4, 4}, {6, 2, 10}, {6, 8, 6}, {7, 9, 4}, {1, 5, 4}, {1, 0, 4}, {9, 7, 3}, {7, 0, 5}, {6, 5, 8}, {1, 7, 6}, {4, 0, 9}, {5, 9, 1}, {8, 7, 3}, {1, 2, 6}, {4, 1, 5}, {5, 2, 4}, {1, 9, 1}, {7, 8, 10}, {0, 4, 2}, {7, 2, 8}},
			src:     6, dst: 0, k: 7, price: 14,
		},
		{
			n:       5,
			flights: [][]int{{0, 1, 5}, {1, 2, 5}, {0, 3, 2}, {3, 1, 2}, {1, 4, 1}, {4, 2, 1}},
			src:     0, dst: 2, k: 2, price: 7,
		},
	} {
		price := findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.k)
		assert.Equal(t, tc.price, price)
	}
}
