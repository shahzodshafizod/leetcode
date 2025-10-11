package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberOfArrays$
func TestNumberOfArrays(t *testing.T) {
	for _, tc := range []struct {
		s   string
		k   int
		cnt int
	}{
		{s: "1000", k: 10000, cnt: 1},
		{s: "1000", k: 10, cnt: 0},
		{s: "1317", k: 2000, cnt: 8},
		{s: "1111111111111", k: 1000000000, cnt: 4076},
	} {
		cnt := numberOfArrays(tc.s, tc.k)
		assert.Equal(t, tc.cnt, cnt)
	}
}
