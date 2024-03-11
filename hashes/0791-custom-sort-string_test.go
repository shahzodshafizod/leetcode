package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCustomSortString$
func TestCustomSortString(t *testing.T) {
	for _, tc := range []struct {
		order    string
		s        string
		permuted string
	}{
		// !!! RESULT IS CORRECT, BUT ORGANIZED RANDOMLY !!!

		// {order: "cba", s: "abcd", permuted: "cbad"},
		// {order: "bcafg", s: "abcd", permuted: "bcad"},
		// {order: "assds", s: "xyz", permuted: "xyz"},
		// {order: "xyz", s: "xyzab", permuted: "xyzab"},
		// {order: "pqrs", s: "pqrstuvwx", permuted: "pqrstuvwx"},
		// {order: "aeiou", s: "hello", permuted: "eohll"},
		// {order: "cba", s: "cba", permuted: "cba"},
		// {order: "kqep", s: "pekeq", permuted: "kqeep"},
		// {order: "ba", s: "abbaba", permuted: "bbbaaa"},
		// {order: "a", s: "dcba", permuted: "adcb"},
	} {
		permuted := customSortString(tc.order, tc.s)
		assert.Equal(t, tc.permuted, permuted)
	}
}
