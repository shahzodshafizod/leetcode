package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSmallestSufficientTeam$
func TestSmallestSufficientTeam(t *testing.T) {
	for _, tc := range []struct {
		reqSkills []string
		people    [][]string
		team      []int
	}{
		{reqSkills: []string{"java", "nodejs", "reactjs"}, people: [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}, team: []int{0, 2}},
		{reqSkills: []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}, people: [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}}, team: []int{1, 2}},
		{reqSkills: []string{"uhppib", "mgdipkgt", "vaostwmycy", "bjwxnfbbubpnd"}, people: [][]string{{"vaostwmycy"}, {"mgdipkgt"}, {"bjwxnfbbubpnd"}, {}, {"uhppib"}}, team: []int{0, 1, 2, 4}},
	} {
		team := smallestSufficientTeam(tc.reqSkills, tc.people)
		assert.Equal(t, tc.team, team)
	}
}
