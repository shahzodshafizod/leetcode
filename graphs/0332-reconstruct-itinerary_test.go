package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestFindItinerary$
func TestFindItinerary(t *testing.T) {
	for _, tc := range []struct {
		tickets     [][]string
		itineraries []string
	}{
		{tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, itineraries: []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, itineraries: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
		{tickets: [][]string{{"JFK", "NRT"}, {"JFK", "KUL"}, {"NRT", "JFK"}}, itineraries: []string{"JFK", "NRT", "JFK", "KUL"}},
		{tickets: [][]string{{"JFK", "ATL"}, {"ORD", "PHL"}, {"JFK", "ORD"}, {"PHX", "LAX"}, {"LAX", "JFK"}, {"PHL", "ATL"}, {"ATL", "PHX"}}, itineraries: []string{"JFK", "ATL", "PHX", "LAX", "JFK", "ORD", "PHL", "ATL"}},
		{tickets: [][]string{{"EZE", "TIA"}, {"EZE", "HBA"}, {"AXA", "TIA"}, {"JFK", "AXA"}, {"ANU", "JFK"}, {"ADL", "ANU"}, {"TIA", "AUA"}, {"ANU", "AUA"}, {"ADL", "EZE"}, {"ADL", "EZE"}, {"EZE", "ADL"}, {"AXA", "EZE"}, {"AUA", "AXA"}, {"JFK", "AXA"}, {"AXA", "AUA"}, {"AUA", "ADL"}, {"ANU", "EZE"}, {"TIA", "ADL"}, {"EZE", "ANU"}, {"AUA", "ANU"}}, itineraries: []string{"JFK", "AXA", "AUA", "ADL", "ANU", "AUA", "ANU", "EZE", "ADL", "EZE", "ANU", "JFK", "AXA", "EZE", "TIA", "AUA", "AXA", "TIA", "ADL", "EZE", "HBA"}},
		{tickets: [][]string{{"MEL", "PER"}, {"SYD", "CBR"}, {"AUA", "DRW"}, {"JFK", "EZE"}, {"PER", "AXA"}, {"DRW", "AUA"}, {"EZE", "SYD"}, {"AUA", "MEL"}, {"DRW", "AUA"}, {"PER", "ANU"}, {"CBR", "EZE"}, {"EZE", "PER"}, {"MEL", "EZE"}, {"EZE", "MEL"}, {"EZE", "TBI"}, {"ADL", "DRW"}, {"ANU", "EZE"}, {"AXA", "ADL"}}, itineraries: []string{"JFK", "EZE", "MEL", "EZE", "PER", "AXA", "ADL", "DRW", "AUA", "DRW", "AUA", "MEL", "PER", "ANU", "EZE", "SYD", "CBR", "EZE", "TBI"}},
		{tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}}, itineraries: []string{"JFK", "SFO", "JFK", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL", "AAA", "BBB", "ATL"}},
		{tickets: [][]string{{"JFK", "ATL"}, {"ATL", "JFK"}}, itineraries: []string{"JFK", "ATL", "JFK"}},
		{tickets: [][]string{{"INN", "CBR"}, {"AXA", "ADL"}, {"HBA", "SYD"}, {"CBR", "INN"}, {"AUA", "MEL"}, {"CBR", "ADL"}, {"PER", "LST"}, {"OOL", "HBA"}, {"LST", "SYD"}, {"EZE", "LST"}, {"HBA", "EZE"}, {"MEL", "LST"}, {"INN", "CNS"}, {"ADL", "INN"}, {"BNE", "MEL"}, {"AXA", "CBR"}, {"CNS", "OOL"}, {"LST", "AXA"}, {"AXA", "CBR"}, {"HBA", "CBR"}, {"JFK", "EZE"}, {"AUA", "EZE"}, {"CBR", "AXA"}, {"LST", "BNE"}, {"AUA", "BIM"}, {"AUA", "MEL"}, {"CBR", "HBA"}, {"AXA", "LST"}, {"PER", "ANU"}, {"ANU", "AUA"}, {"SYD", "PER"}, {"MEL", "PER"}, {"CBR", "MEL"}, {"BNE", "AUA"}, {"LST", "CBR"}, {"EZE", "AUA"}, {"LST", "PER"}, {"MEL", "BNE"}, {"PER", "AXA"}, {"EZE", "HBA"}, {"MEL", "AUA"}, {"SYD", "AXA"}, {"ADL", "LST"}}, itineraries: []string{"JFK", "EZE", "AUA", "EZE", "HBA", "CBR", "ADL", "INN", "CBR", "AXA", "ADL", "LST", "AXA", "CBR", "HBA", "EZE", "LST", "BNE", "AUA", "MEL", "AUA", "MEL", "BNE", "MEL", "LST", "CBR", "INN", "CNS", "OOL", "HBA", "SYD", "AXA", "CBR", "MEL", "PER", "AXA", "LST", "PER", "LST", "SYD", "PER", "ANU", "AUA", "BIM"}},
	} {
		itineraries := findItinerary(tc.tickets)
		assert.Equal(t, tc.itineraries, itineraries)
	}
}
