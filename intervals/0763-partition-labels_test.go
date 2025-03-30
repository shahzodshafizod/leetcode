package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestPartitionLabels$
func TestPartitionLabels(t *testing.T) {
	for _, tc := range []struct {
		s     string
		sizes []int
	}{
		{s: "ababcbacadefegdehijhklij", sizes: []int{9, 7, 8}},
		{s: "eccbbbbdec", sizes: []int{10}},
	} {
		sizes := partitionLabels(tc.s)
		assert.Equal(t, tc.sizes, sizes)
	}
}
