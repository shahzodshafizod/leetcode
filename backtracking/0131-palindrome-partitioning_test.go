package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestPartition$
func TestPartition(t *testing.T) {
	for _, tc := range []struct {
		s             string
		partitionings [][]string
	}{
		{s: "aab", partitionings: [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{s: "a", partitionings: [][]string{{"a"}}},
		{s: "efe", partitionings: [][]string{{"e", "f", "e"}, {"efe"}}},
	} {
		partitionings := partition(tc.s)
		assert.Equal(t, tc.partitionings, partitionings)
	}
}
