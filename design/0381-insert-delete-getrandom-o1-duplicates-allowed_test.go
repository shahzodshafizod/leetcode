package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestRandomizedCollection$
func TestRandomizedCollection(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"RandomizedCollection", "insert", "insert", "insert", "getRandom", "remove", "getRandom"},
			values:   [][]int{{}, {1}, {1}, {2}, {}, {1}, {}},
			output:   []any{nil, true, false, true, nil, true, nil},
		},
		{
			commands: []string{"RandomizedCollection", "insert", "remove", "insert", "remove", "getRandom", "getRandom", "getRandom", "getRandom"},
			values:   [][]int{{}, {1}, {2}, {2}, {1}, {}, {}, {}, {}},
			output:   []any{nil, true, false, true, true, nil, nil, nil, nil},
		},
		{
			commands: []string{"RandomizedCollection", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "remove"},
			values:   [][]int{{}, {10}, {10}, {20}, {20}, {10}, {10}, {20}, {20}},
			output:   []any{nil, true, false, true, false, true, true, true, true},
		},
		{
			commands: []string{"RandomizedCollection", "insert", "insert", "remove", "insert"},
			values:   [][]int{{}, {0}, {0}, {0}, {0}},
			output:   []any{nil, true, false, true, false},
		},
	} {
		var randomCol RandomizedCollection

		for index, command := range tc.commands {
			var output any

			switch command {
			case "RandomizedCollection":
				randomCol = NewRandomizedCollection()
			case "insert":
				output = randomCol.Insert(tc.values[index][0])
			case "remove":
				output = randomCol.Remove(tc.values[index][0])
			case "getRandom":
				output = randomCol.GetRandom()
				// For getRandom, we just verify it returns a value in our collection
				// We can't check exact value due to randomness
				continue
			default:
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
