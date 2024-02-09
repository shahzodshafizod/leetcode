package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestPredictPartyVictory$
func TestPredictPartyVictory(t *testing.T) {
	for _, tc := range []struct {
		senate string
		party  string
	}{
		{senate: "RD", party: "Radiant"},
		{senate: "RDD", party: "Dire"},
		{senate: "DDRRR", party: "Dire"},
		{senate: "DRDRR", party: "Dire"},
		{senate: "DRRDRDRDRDDRDRDR", party: "Radiant"},
		{senate: "RRDDDDDDDRRDRRDDRRRR", party: "Dire"},
	} {
		party := predictPartyVictory(tc.senate)
		assert.Equal(t, tc.party, party)
	}
}
