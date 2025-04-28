package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestUncommonFromSentences$
func TestUncommonFromSentences(t *testing.T) {
	for _, tc := range []struct {
		s1       string
		s2       string
		uncommon []string
	}{
		{s1: "this apple is sweet", s2: "this apple is sour", uncommon: []string{"sweet", "sour"}},
		{s1: "apple apple", s2: "banana", uncommon: []string{"banana"}},
		{s1: "this apple is sweet sweet sweet", s2: "this apple is sour", uncommon: []string{"sour"}},
		{s1: "this apple is sweet sweet sweet and good", s2: "this apple is sour", uncommon: []string{"and", "good", "sour"}},
		{s1: "s z z z s", s2: "s z ejt", uncommon: []string{"ejt"}},
	} {
		uncommon := uncommonFromSentences(tc.s1, tc.s2)
		assert.ElementsMatch(t, tc.uncommon, uncommon)
	}
}
