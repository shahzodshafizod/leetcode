package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestShortestCompletingWord$
func TestShortestCompletingWord(t *testing.T) {
	for _, tc := range []struct {
		licensePlate string
		words        []string
		expected     string
	}{
		{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}, "steps"},
		{"1s3 456", []string{"looks", "pest", "stew", "show"}, "pest"},
		{"Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}, "husband"},
		{"OgEu755", []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}, "enough"},
	} {
		output := shortestCompletingWord(tc.licensePlate, tc.words)
		assert.Equal(t, tc.expected, output)
	}
}
