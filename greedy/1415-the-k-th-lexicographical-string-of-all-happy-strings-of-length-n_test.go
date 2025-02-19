package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestGetHappyString$
func TestGetHappyString(t *testing.T) {
	for _, tc := range []struct {
		n     int
		k     int
		happy string
	}{
		{n: 1, k: 3, happy: "c"},
		{n: 1, k: 4, happy: ""},
		{n: 3, k: 9, happy: "cab"},
	} {
		happy := getHappyString(tc.n, tc.k)
		assert.Equal(t, tc.happy, happy)
	}
}
