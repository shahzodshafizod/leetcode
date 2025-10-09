package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestSmallestSufficientTeam$
func TestSmallestSufficientTeam(t *testing.T) {
	for _, tc := range []struct {
		req_skills []string
		people     [][]string
		team       []int
	}{
		{req_skills: []string{"java", "nodejs", "reactjs"}, people: [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}, team: []int{0, 2}},
		{req_skills: []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"}, people: [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}}, team: []int{1, 2}},
		{req_skills: []string{"uhppib", "mgdipkgt", "vaostwmycy", "bjwxnfbbubpnd"}, people: [][]string{{"vaostwmycy"}, {"mgdipkgt"}, {"bjwxnfbbubpnd"}, {}, {"uhppib"}}, team: []int{0, 1, 2, 4}},
	} {
		team := smallestSufficientTeam(tc.req_skills, tc.people)
		assert.Equal(t, tc.team, team)
	}
}
