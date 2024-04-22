package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestOpenLock$
func TestOpenLock(t *testing.T) {
	for _, tc := range []struct {
		deadends []string
		target   string
		turns    int
	}{
		{deadends: []string{"0201", "0101", "0102", "1212", "2002"}, target: "0202", turns: 6},
		{deadends: []string{"8888"}, target: "0009", turns: 1},
		{deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, target: "8888", turns: -1},
		{deadends: []string{"1002", "1220", "0122", "0112", "0121"}, target: "1200", turns: 3},
	} {
		turns := openLock(tc.deadends, tc.target)
		assert.Equal(t, tc.turns, turns)
	}
}
