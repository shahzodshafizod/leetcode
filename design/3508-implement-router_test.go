package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestRouter$
func TestRouter(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   [][]int
		output   []any
	}{
		{
			commands: []string{"Router", "addPacket", "addPacket", "addPacket", "addPacket", "addPacket", "forwardPacket", "addPacket", "getCount"},
			values:   [][]int{{3}, {1, 4, 90}, {2, 5, 90}, {1, 4, 90}, {3, 5, 95}, {4, 5, 105}, {}, {5, 2, 110}, {5, 100, 110}},
			output:   []any{nil, true, true, false, true, true, []int{2, 5, 90}, true, 1},
		},
		{
			commands: []string{"Router", "addPacket", "forwardPacket", "forwardPacket"},
			values:   [][]int{{2}, {7, 4, 90}, {}, {}},
			output:   []any{nil, true, []int{7, 4, 90}, []int{}},
		},
		{
			commands: []string{"Router", "addPacket", "addPacket", "addPacket", "getCount", "getCount", "addPacket", "forwardPacket", "addPacket"},
			values:   [][]int{{5}, {2, 3, 1}, {5, 2, 5}, {2, 3, 5}, {3, 4, 4}, {3, 5, 5}, {3, 2, 5}, {}, {2, 3, 5}},
			output:   []any{nil, true, true, true, 0, 1, true, []int{2, 3, 1}, false},
		},
	} {
		var router Router

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "Router":
				router = NewRouter(tc.values[idx][0])
			case "addPacket":
				output = router.AddPacket(tc.values[idx][0], tc.values[idx][1], tc.values[idx][2])
			case "forwardPacket":
				output = router.ForwardPacket()
			case "getCount":
				output = router.GetCount(tc.values[idx][0], tc.values[idx][1], tc.values[idx][2])
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
