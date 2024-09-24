package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestKthSmallestPath$
func TestKthSmallestPath(t *testing.T) {
	for _, tc := range []struct {
		destination []int
		k           int
		path        string
	}{
		{destination: []int{2, 3}, k: 1, path: "HHHVV"},
		{destination: []int{2, 3}, k: 2, path: "HHVHV"},
		{destination: []int{2, 3}, k: 3, path: "HHVVH"},
		{destination: []int{15, 15}, k: 155117520, path: "VVVVVVVVVVVVVVVHHHHHHHHHHHHHHH"},
	} {
		path := kthSmallestPath(tc.destination, tc.k)
		assert.Equal(t, tc.path, path)
	}
}
