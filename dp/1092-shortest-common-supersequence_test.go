package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestShortestCommonSupersequence$
func TestShortestCommonSupersequence(t *testing.T) {
	for _, tc := range []struct {
		str1 string
		str2 string
		csc  string
	}{
		{str1: "abac", str2: "cab", csc: "cabac"},
		{str1: "aaaaaaaa", str2: "aaaaaaaa", csc: "aaaaaaaa"},
	} {
		csc := shortestCommonSupersequence(tc.str1, tc.str2)
		assert.Equal(t, tc.csc, csc)
	}
}
