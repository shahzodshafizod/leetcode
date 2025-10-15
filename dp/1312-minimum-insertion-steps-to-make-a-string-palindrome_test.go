package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinInsertions$
func TestMinInsertions(t *testing.T) {
	for _, tc := range []struct {
		s   string
		cnt int
	}{
		{s: "zzazz", cnt: 0},
		{s: "mbadm", cnt: 2},
		{s: "leetcode", cnt: 5},
	} {
		cnt := minInsertions(tc.s)
		assert.Equal(t, tc.cnt, cnt)
	}
}
