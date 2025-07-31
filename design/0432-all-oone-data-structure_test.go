package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestAllOne$
func TestAllOne(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]string
		output   []any
	}{
		{
			commands: []string{"AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"},
			values:   [][]string{{}, {"hello"}, {"hello"}, {}, {}, {"leet"}, {}, {}},
			output:   []any{nil, nil, nil, "hello", "hello", nil, "hello", "leet"},
		},
		{
			commands: []string{"AllOne", "inc", "inc", "inc", "inc", "getMaxKey", "inc", "inc", "inc", "dec", "inc", "inc", "inc", "getMaxKey"},
			values:   [][]string{{}, {"hello"}, {"goodbye"}, {"hello"}, {"hello"}, {}, {"leet"}, {"code"}, {"leet"}, {"hello"}, {"leet"}, {"code"}, {"code"}, {}},
			output:   []any{nil, nil, nil, nil, nil, "hello", nil, nil, nil, nil, nil, nil, nil, "leet"},
		},
		{
			commands: []string{"AllOne", "inc", "inc", "inc", "inc", "inc", "inc", "dec", "dec", "getMinKey", "dec", "getMaxKey", "getMinKey"},
			values:   [][]string{{}, {"a"}, {"b"}, {"b"}, {"c"}, {"c"}, {"c"}, {"b"}, {"b"}, {}, {"a"}, {}, {}},
			output:   []any{nil, nil, nil, nil, nil, nil, nil, nil, nil, "a", nil, "c", "c"},
		},
	} {
		var allOne AllOne

		for index, command := range tc.commands {
			var output any

			switch command {
			case "AllOne":
				allOne = NewAllOne()
			case "inc":
				allOne.Inc(tc.values[index][0])
			case "dec":
				allOne.Dec(tc.values[index][0])
			case "getMaxKey":
				output = allOne.GetMaxKey()
			case "getMinKey":
				output = allOne.GetMinKey()
			}

			assert.Equal(t, tc.output[index], output)
		}
	}
}
