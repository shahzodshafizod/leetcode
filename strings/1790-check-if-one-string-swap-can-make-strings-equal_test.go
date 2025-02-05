package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestAreAlmostEqual$
func TestAreAlmostEqual(t *testing.T) {
	for _, tc := range []struct {
		s1    string
		s2    string
		equal bool
	}{
		{s1: "bank", s2: "kanb", equal: true},
		{s1: "attack", s2: "defend", equal: false},
		{s1: "kelb", s2: "kelb", equal: true},
		{s1: "xkkkkkkkx", s2: "fkkkkkkkf", equal: false},
	} {
		equal := areAlmostEqual(tc.s1, tc.s2)
		assert.Equal(t, tc.equal, equal)
	}
}
