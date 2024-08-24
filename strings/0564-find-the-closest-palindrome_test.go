package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestNearestPalindromic$
func TestNearestPalindromic(t *testing.T) {
	for _, tc := range []struct {
		n       string
		nearest string
	}{
		{n: "123", nearest: "121"},
		{n: "1", nearest: "0"},
		{n: "139", nearest: "141"},
		{n: "1234", nearest: "1221"},
		{n: "999", nearest: "1001"},
		{n: "1000", nearest: "999"},
		{n: "12932", nearest: "12921"},
		{n: "99800", nearest: "99799"},
		{n: "12120", nearest: "12121"},
		{n: "99899", nearest: "99799"},
		{n: "101", nearest: "99"},
		{n: "11", nearest: "9"},
		{n: "12345", nearest: "12321"},
		{n: "10", nearest: "9"},
		{n: "88", nearest: "77"},
		{n: "1837722381", nearest: "1837667381"},
		{n: "1000000000000000", nearest: "999999999999999"},
		{n: "999999999999999999", nearest: "1000000000000000001"},
		{n: "807045053224792883", nearest: "807045053350540708"},
	} {
		nearest := nearestPalindromic(tc.n)
		assert.Equal(t, tc.nearest, nearest)
	}
}
