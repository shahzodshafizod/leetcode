package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMovieRentingSystem$
func TestMovieRentingSystem(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   []any
		output   []any
	}{
		{
			commands: []string{"MovieRentingSystem", "search", "rent", "rent", "report", "drop", "search"},
			values:   []any{[]any{3, [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}}}, []int{1}, []int{0, 1}, []int{1, 2}, []int{}, []int{1, 2}, []int{2}},
			output:   []any{nil, []int{1, 0, 2}, nil, nil, [][]int{{0, 1}, {1, 2}}, nil, []int{0, 1}},
		},
		{
			commands: []string{"MovieRentingSystem", "rent", "drop", "rent", "rent", "rent", "drop", "search", "report", "rent", "search"},
			values:   []any{[]any{10, [][]int{{4, 374, 55}, {1, 6371, 21}, {8, 3660, 24}, {1, 56, 32}, {5, 374, 71}, {3, 4408, 36}, {6, 9322, 73}, {6, 9574, 92}, {8, 7834, 62}, {2, 6084, 27}, {7, 3262, 89}, {2, 8959, 53}, {0, 3323, 41}, {6, 6565, 45}, {0, 4239, 20}}}, []int{0, 4239}, []int{0, 4239}, []int{3, 4408}, []int{2, 6084}, []int{0, 4239}, []int{0, 4239}, []int{9346}, []int{}, []int{6, 9322}, []int{8698}},
			output:   []any{nil, nil, nil, nil, nil, nil, nil, []int{}, [][]int{{2, 6084}, {3, 4408}}, nil, []int{}},
		},
		{
			commands: []string{"MovieRentingSystem", "search", "search", "rent", "search", "search", "report", "search", "drop"},
			values:   []any{[]any{10, [][]int{{0, 418, 3}, {9, 5144, 46}, {2, 8986, 29}, {6, 1446, 28}, {3, 8215, 97}, {7, 9105, 34}, {6, 9105, 30}, {5, 1722, 94}, {9, 528, 40}, {3, 850, 77}, {3, 7069, 40}, {8, 1997, 42}, {0, 8215, 28}, {7, 4050, 80}, {4, 7100, 97}, {4, 9686, 32}, {4, 2566, 93}, {2, 8320, 12}, {2, 5495, 56}}}, []int{7837}, []int{5495}, []int{4, 7100}, []int{9105}, []int{1446}, []int{}, []int{9869}, []int{4, 7100}},
			output:   []any{nil, []int{}, []int{2}, nil, []int{6, 7}, []int{6}, [][]int{{4, 7100}}, []int{}, nil},
		},
	} {
		var renting MovieRentingSystem

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "MovieRentingSystem":
				values, ok := tc.values[idx].([]any)
				_ = ok
				n, ok := values[0].(int)
				_ = ok
				entries, ok := values[1].([][]int)
				_ = ok
				renting = NewMovieRentingSystem(n, entries)
			case "search":
				values, ok := tc.values[idx].([]int)
				_ = ok
				output = renting.Search(values[0])
			case "rent":
				values, ok := tc.values[idx].([]int)
				_ = ok

				renting.Rent(values[0], values[1])
			case "drop":
				values, ok := tc.values[idx].([]int)
				_ = ok

				renting.Drop(values[0], values[1])
			case "report":
				output = renting.Report()
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
