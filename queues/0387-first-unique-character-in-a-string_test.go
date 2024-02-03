package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestFirstUniqChar$
func TestFirstUniqChar(t *testing.T) {
	for _, tc := range []struct {
		s     string
		index int
	}{
		{s: "leetcode", index: 0},
		{s: "loveleetcode", index: 2},
		{s: "aabb", index: -1},
		{s: "itwqbtcdprfsuprkrjkausiterybzncbmdvkgljxuekizvaivszowqtmrttiihervpncztuoljftlxybpgwnjb", index: 61},
	} {
		index := firstUniqChar(tc.s)
		assert.Equal(t, tc.index, index)
	}
}
