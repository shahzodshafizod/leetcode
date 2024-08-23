package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestGetPermutation$
func TestGetPermutation(t *testing.T) {
	for _, tc := range []struct {
		n           int
		k           int
		permutation string
	}{
		{n: 3, k: 3, permutation: "213"},
		{n: 4, k: 9, permutation: "2314"},
		{n: 3, k: 1, permutation: "123"},
		{n: 4, k: 17, permutation: "3412"},
	} {
		permutation := getPermutation(tc.n, tc.k)
		assert.Equal(t, tc.permutation, permutation)
	}
}
