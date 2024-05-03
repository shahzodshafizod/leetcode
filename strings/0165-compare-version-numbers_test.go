package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestCompareVersion$
func TestCompareVersion(t *testing.T) {
	for _, tc := range []struct {
		version1 string
		version2 string
		result   int
	}{
		{version1: "1.01", version2: "1.001", result: 0},
		{version1: "1.01", version2: "1.1", result: 0},
		{version1: "1.0", version2: "1.0.0", result: 0},
		{version1: "0.1", version2: "1.1", result: -1},
		{version1: "3.3.0", version2: "3.3", result: 0},
		{version1: "1.2", version2: "1.10", result: -1},
		{version1: "1", version2: "0", result: 1},
		{version1: "1.05", version2: "1.1", result: 1},
		{version1: "1.1", version2: "1.10", result: -1},
		{version1: "2.10000000000000000000000000000.0", version2: "2.1", result: 1},
		{version1: "2.105", version2: "2.22", result: 1},
	} {
		result := compareVersion(tc.version1, tc.version2)
		assert.Equal(t, tc.result, result)
	}
}
