package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestNumberOfWays$
func TestNumberOfWays(t *testing.T) {
	for _, tc := range []struct {
		corridor string
		ways     int
	}{
		{corridor: "SSPPSPS", ways: 3},
		{corridor: "PPSPSP", ways: 1},
		{corridor: "S", ways: 0},
		{corridor: "P", ways: 0},
		{corridor: "PPSPSPPSPSPSPSPS", ways: 0},
		{corridor: "SPPSSSSPPS", ways: 1},
		{corridor: "SPSPPSSPSSSS", ways: 6},
		{corridor: "PPPSPPPSPSSPPSPSSPSSPPPPSSPSSPPSPPPSSSPSSSPSSSSPSSSSSPSSPSPPSSPSSPPSSSPSPPPSSSSSPSSPPPSSPPSSPSSSPPSP", ways: 0},
		{corridor: "PPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPP", ways: 0},
		{corridor: "PSPSPPPSSSSPSSPSSPSPPSSSPPSPSPPPSSSPSPPSSSPSSSPPSPPPSPSPPPSPSPPSSPPSSPSPSSPPPPPSPSPSSSPSSSSPPPSPSPPS", ways: 663552},
		{corridor: "SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS", ways: 1},
	} {
		ways := numberOfWays(tc.corridor)
		assert.Equal(t, tc.ways, ways)
	}
}
