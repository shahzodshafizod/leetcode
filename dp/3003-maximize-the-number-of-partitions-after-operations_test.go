package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxPartitionsAfterOperations$
func TestMaxPartitionsAfterOperations(t *testing.T) {
	for _, tc := range []struct {
		s          string
		k          int
		partitions int
	}{
		{s: "accca", k: 2, partitions: 3},
		{s: "aabaab", k: 3, partitions: 1},
		{s: "xxyz", k: 1, partitions: 4},
	} {
		partitions := maxPartitionsAfterOperations(tc.s, tc.k)
		assert.Equal(t, tc.partitions, partitions)
	}
}
