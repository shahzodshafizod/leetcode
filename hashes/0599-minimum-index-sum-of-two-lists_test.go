package hashes

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindRestaurant$
func TestFindRestaurant(t *testing.T) {
	for _, tc := range []struct {
		list1  []string
		list2  []string
		output []string
	}{
		{
			list1:  []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:  []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			output: []string{"Shogun"},
		},
		{
			list1:  []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:  []string{"KFC", "Shogun", "Burger King"},
			output: []string{"Shogun"},
		},
		{
			list1:  []string{"happy", "sad", "good"},
			list2:  []string{"sad", "happy", "good"},
			output: []string{"sad", "happy"},
		},
	} {
		output := findRestaurant(tc.list1, tc.list2)
		// Sort both for comparison since order doesn't matter
		sort.Strings(output)
		sort.Strings(tc.output)
		assert.Equal(t, tc.output, output)
	}
}
