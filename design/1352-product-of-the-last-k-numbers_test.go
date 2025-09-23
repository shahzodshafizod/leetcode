package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestProductOfNumbers$
func TestProductOfNumbers(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"},
			values:   [][]int{{}, {3}, {0}, {2}, {5}, {4}, {2}, {3}, {4}, {8}, {2}},
			output:   []any{nil, nil, nil, nil, nil, nil, 20, 40, 0, nil, 32},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "getProduct", "add", "add", "add", "add", "add", "getProduct", "add", "getProduct", "add", "getProduct", "getProduct", "add", "add", "add", "add"},
			values:   [][]int{{}, {2}, {2}, {5}, {10}, {2}, {0}, {5}, {0}, {2}, {3}, {3}, {10}, {4}, {7}, {2}, {1}, {8}, {2}, {6}, {5}},
			output:   []any{nil, nil, nil, nil, nil, 50, nil, nil, nil, nil, nil, 0, nil, 0, nil, 70, 7, nil, nil, nil, nil},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "add", "add", "getProduct"},
			values:   [][]int{{}, {1}, {1}, {0}, {5}, {4}, {1}, {1}, {0}, {3}},
			output:   []any{nil, nil, nil, nil, nil, nil, 4, nil, nil, 0},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct"},
			values:   [][]int{{}, {9}, {4}, {10}, {5}, {0}, {1}},
			output:   []any{nil, nil, nil, nil, nil, nil, 0},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "getProduct", "getProduct", "add", "getProduct", "getProduct", "add", "add", "add", "add", "add", "getProduct", "getProduct", "add", "getProduct"},
			values:   [][]int{{}, {2}, {7}, {10}, {9}, {3}, {1}, {5}, {1}, {1}, {6}, {10}, {9}, {2}, {9}, {2}, {4}, {1}, {3}},
			output:   []any{nil, nil, nil, nil, nil, 630, 9, nil, 5, 5, nil, nil, nil, nil, nil, 18, 1620, nil, 18},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "getProduct", "getProduct"},
			values:   [][]int{{}, {7}, {8}, {0}, {1}, {1}, {2}},
			output:   []any{nil, nil, nil, nil, nil, 1, 0},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"},
			values:   [][]int{{}, {8}, {2}, {7}, {7}, {3}, {7}, {4}, {4}, {4}, {0}, {2}},
			output:   []any{nil, nil, nil, nil, nil, nil, nil, 1029, 1029, 1029, nil, 0},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "getProduct", "add", "getProduct", "add", "add", "add", "getProduct", "add", "getProduct", "add", "getProduct", "add", "getProduct", "getProduct", "getProduct", "getProduct"},
			values:   [][]int{{}, {3}, {7}, {3}, {6}, {2}, {0}, {3}, {3}, {4}, {7}, {3}, {3}, {3}, {8}, {3}, {8}, {4}, {2}, {4}, {4}},
			output:   []any{nil, nil, nil, nil, nil, 18, nil, 0, nil, nil, nil, 84, nil, 84, nil, 168, nil, 1344, 64, 1344, 1344},
		},
		{
			commands: []string{"ProductOfNumbers", "add", "add", "add", "add", "add"},
			values:   [][]int{{}, {0}, {6}, {7}, {5}, {5}},
			output:   []any{nil, nil, nil, nil, nil, nil},
		},
	} {
		var p ProductOfNumbers

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "ProductOfNumbers":
				p = NewProductOfNumbers()
			case "add":
				p.Add(tc.values[idx][0])
			case "getProduct":
				output = p.GetProduct(tc.values[idx][0])
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
