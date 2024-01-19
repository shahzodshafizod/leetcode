package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestBackspaceCompare$
func TestBackspaceCompare(t *testing.T) {
	for _, tc := range []struct {
		s     string
		t     string
		equal bool
	}{
		{s: "a###b", t: "", equal: false},
		{s: "ab#z", t: "az#z", equal: true},
		{s: "abc#d", t: "acc#d", equal: false},
		{s: "x#y#z#", t: "a#", equal: true},
		{s: "ax#y#z#", t: "a", equal: true},
		{s: "a###b", t: "b", equal: true},
		{s: "Ab#z", t: "ab#z", equal: false},
		{s: "abc", t: "abc", equal: true},
		{s: "a########", t: "", equal: true},
		{s: "ab#c", t: "ad#c", equal: true},
		{s: "ab##", t: "c#d#", equal: true},
		{s: "a#c", t: "b", equal: false},
		{s: "bxj##tw", t: "bxo#j##tw", equal: true},
	} {
		equal := backspaceCompare(tc.s, tc.t)
		assert.Equal(t, tc.equal, equal)
	}
}
