package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestFullJustify$
func TestFullJustify(t *testing.T) {
	for _, tc := range []struct {
		words    []string
		maxWidth int
		lines    []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			lines:    []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			lines:    []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			words:    []string{"Here", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 14,
			lines:    []string{"Here   is   an", "example     of", "text          ", "justification."},
		},
		{
			words:    []string{"What", "must", "be", "shall", "be."},
			maxWidth: 12,
			lines:    []string{"What must be", "shall be.   "},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			lines:    []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
		{
			words:    []string{"My", "momma", "always", "said,", "Life", "was", "like", "a", "box", "of", "chocolates.", "You", "never", "know", "what", "you're", "gonna", "get."},
			maxWidth: 20,
			lines:    []string{"My    momma   always", "said,  Life was like", "a box of chocolates.", "You  never know what", "you're gonna get.   "},
		},
		{
			words:    []string{"a", "b", "c", "d"},
			maxWidth: 3,
			lines:    []string{"a b", "c d"},
		},
		{
			words:    []string{"a"},
			maxWidth: 2,
			lines:    []string{"a "},
		},
		{
			words:    []string{"", "", "", "", "a"},
			maxWidth: 4,
			lines:    []string{"    ", "a   "},
		},
	} {
		lines := fullJustify(tc.words, tc.maxWidth)
		assert.Equal(t, tc.lines, lines)
	}
}
