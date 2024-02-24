package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestFindAllPeople$
func TestFindAllPeople(t *testing.T) {
	for _, tc := range []struct {
		n           int
		meetings    [][]int
		firstPerson int
		informed    []int
	}{
		{n: 6, meetings: [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, firstPerson: 1, informed: []int{0, 1, 2, 3, 5}},
		{n: 4, meetings: [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, firstPerson: 3, informed: []int{0, 1, 3}},
		{n: 5, meetings: [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, firstPerson: 1, informed: []int{0, 1, 2, 3, 4}},
		{n: 5, meetings: [][]int{{1, 4, 3}, {0, 4, 3}}, firstPerson: 3, informed: []int{0, 1, 3, 4}},
	} {
		informed := findAllPeople(tc.n, tc.meetings, tc.firstPerson)
		assert.Equal(t, tc.informed, informed)
	}
}
