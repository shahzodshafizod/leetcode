package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestBeautifulIndices$
func TestBeautifulIndices(t *testing.T) {
	for _, tc := range []struct {
		s        string
		a        string
		b        string
		k        int
		expected []int
	}{
		{
			s:        "isawsquirrelnearmysquirrelhouseohmy",
			a:        "my",
			b:        "squirrel",
			k:        15,
			expected: []int{16, 33},
		},
		{
			s:        "abcd",
			a:        "a",
			b:        "a",
			k:        4,
			expected: []int{0},
		},
		{
			s:        "lahhnlwx",
			a:        "hhnlw",
			b:        "ty",
			k:        6,
			expected: nil,
		},
		{
			s:        "sqgxo",
			a:        "s",
			b:        "xo",
			k:        3,
			expected: []int{0},
		},
		{
			s:        "aba",
			a:        "a",
			b:        "a",
			k:        1,
			expected: []int{0, 2},
		},
		{
			s:        "anvuyenjdp",
			a:        "jd",
			b:        "nu",
			k:        4,
			expected: nil,
		},
	} {
		output := beautifulIndices(tc.s, tc.a, tc.b, tc.k)
		assert.Equal(t, tc.expected, output)
	}
}
