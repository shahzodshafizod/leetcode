package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestFindRotateSteps$
func TestFindRotateSteps(t *testing.T) {
	for _, tc := range []struct {
		ring  string
		key   string
		steps int
	}{
		{ring: "godding", key: "gd", steps: 4},
		{ring: "godding", key: "godding", steps: 13},
		{ring: "iotfo", key: "fioot", steps: 11},
		{ring: "abccbaxbe", key: "abx", steps: 6},
	} {
		steps := findRotateSteps(tc.ring, tc.key)
		assert.Equal(t, tc.steps, steps)
	}
}
