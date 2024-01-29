package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindJudge$
func TestFindJudge(t *testing.T) {
	for _, tc := range []struct {
		n          int
		trust      [][]int
		judgeLabel int
	}{
		{n: 1, trust: [][]int{}, judgeLabel: 1},
		{n: 2, trust: [][]int{{1, 2}}, judgeLabel: 2},
		{n: 3, trust: [][]int{{1, 3}, {2, 3}}, judgeLabel: 3},
		{n: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}, judgeLabel: -1},
	} {
		judgeLable := findJudge(tc.n, tc.trust)
		assert.Equal(t, tc.judgeLabel, judgeLable)
	}
}
