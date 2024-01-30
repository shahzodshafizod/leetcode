package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestNetworkDelayTime$
func TestNetworkDelayTime(t *testing.T) {
	for _, tc := range []struct {
		times     [][]int
		n         int
		k         int
		delayTime int
	}{
		{times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, n: 4, k: 2, delayTime: 2},
		{times: [][]int{{1, 2, 1}}, n: 2, k: 1, delayTime: 1},
		{times: [][]int{{1, 2, 1}}, n: 2, k: 2, delayTime: -1},
		{times: [][]int{{1, 2, 9}, {1, 4, 2}, {2, 5, 1}, {4, 2, 4}, {4, 5, 6}, {3, 2, 3}, {3, 1, 5}, {5, 3, 7}}, n: 5, k: 1, delayTime: 14},
		{times: [][]int{{3, 5, 78}, {2, 1, 1}, {1, 3, 0}, {4, 3, 59}, {5, 3, 85}, {5, 2, 22}, {2, 4, 23}, {1, 4, 43}, {4, 5, 75}, {5, 1, 15}, {1, 5, 91}, {4, 1, 16}, {3, 2, 98}, {3, 4, 22}, {5, 4, 31}, {1, 2, 0}, {2, 5, 4}, {4, 2, 51}, {3, 1, 36}, {2, 3, 59}}, n: 5, k: 5, delayTime: 31},
		{times: [][]int{{4, 2, 76}, {1, 3, 79}, {3, 1, 81}, {4, 3, 30}, {2, 1, 47}, {1, 5, 61}, {1, 4, 99}, {3, 4, 68}, {3, 5, 46}, {4, 1, 6}, {5, 4, 7}, {5, 3, 44}, {4, 5, 19}, {2, 3, 13}, {3, 2, 18}, {1, 2, 0}, {5, 1, 25}, {2, 5, 58}, {2, 4, 77}, {5, 2, 74}}, n: 5, k: 3, delayTime: 59},
		{times: [][]int{{1, 2, 1}, {2, 3, 7}, {1, 3, 4}, {2, 1, 2}}, n: 4, k: 1, delayTime: -1},
		{times: [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}, n: 3, k: 1, delayTime: 2},
	} {
		delayTime := networkDelayTime(tc.times, tc.n, tc.k)
		assert.Equal(t, tc.delayTime, delayTime)
	}
}
