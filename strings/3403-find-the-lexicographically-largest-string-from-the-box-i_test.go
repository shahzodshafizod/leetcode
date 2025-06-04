package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestAnswerString$
func TestAnswerString(t *testing.T) {
	for _, tc := range []struct {
		word       string
		numFriends int
		largest    string
	}{
		{word: "dbca", numFriends: 2, largest: "dbc"},
		{word: "gggg", numFriends: 4, largest: "g"},
		{word: "aann", numFriends: 2, largest: "nn"},
		{word: "gh", numFriends: 1, largest: "gh"},
		{word: "yxz", numFriends: 3, largest: "z"},
		{word: "xzy", numFriends: 2, largest: "zy"},
		{word: "xxyy", numFriends: 2, largest: "yy"},
		{word: "xy", numFriends: 1, largest: "xy"},
		{word: "xzyz", numFriends: 3, largest: "zy"},
		{word: "adif", numFriends: 1, largest: "adif"},
		{word: "bagj", numFriends: 3, largest: "j"},
		{word: "gm", numFriends: 1, largest: "gm"},
		{word: "gm", numFriends: 2, largest: "m"},
		{word: "gmn", numFriends: 3, largest: "n"},
		{word: "gmhm", numFriends: 3, largest: "mh"},
		{word: "gmng", numFriends: 2, largest: "ng"},
		{word: "gmnaangggg", numFriends: 4, largest: "ngggg"},
	} {
		largest := answerString(tc.word, tc.numFriends)
		assert.Equal(t, tc.largest, largest)
	}
}
