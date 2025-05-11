package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLicenseKeyFormatting$
func TestLicenseKeyFormatting(t *testing.T) {
	for _, tc := range []struct {
		s         string
		k         int
		formatted string
	}{
		{s: "5F3Z-2e-9-w", k: 4, formatted: "5F3Z-2E9W"},
		{s: "2-5g-3-J", k: 2, formatted: "2-5G-3J"},
		{s: "---", k: 3, formatted: ""},
		{s: "2-4A0r7-4k", k: 4, formatted: "24A0-R74K"},
		{s: "5F3Z-2e-9-w", k: 44, formatted: "5F3Z2E9W"},
		{s: "--a-a-a-a--", k: 2, formatted: "AA-AA"},
	} {
		formatted := licenseKeyFormatting(tc.s, tc.k)
		assert.Equal(t, tc.formatted, formatted)
	}
}
