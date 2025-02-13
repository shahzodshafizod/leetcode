package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMinLength$
func TestMinLength(t *testing.T) {
	for _, tc := range []struct {
		s      string
		length int
	}{
		{s: "ABFCACDB", length: 2},
		{s: "ACBBD", length: 5},
		{s: "CCADDADDDBBCDDBABACADABAABADCABDCCBDACBDBAADDABCABBCABBDDAABCBCBBCCCDBCDDDADAACBCACDDBBA", length: 62},
		{s: "DCDCBCBDACBBABDABABCCCBABCCCCCCCBDDBCDACDADABADDCDBBC", length: 35},
		{s: "BBBDCADCDACACDBBCACDACDABCBCDDADCDCACCDDBCACCDDDCCBCDBDCBDDCBCBBCCBBBAACBBB", length: 47},
		{s: "BCDDBCCCCBACCADDCBDADDCCABCCCBACAADDADCDAACABDDDDABBACBABCABCCDCABBA", length: 48},
		{s: "RZAAAACCCCCAABBDDDDDBBBBYAAAAAAACAACAACACCAAAACCACCDDBDDBBBBDDBDBBDBBDBBBBBBBMBSCACCACDBDDBDACCDDBCD", length: 6},
		{s: "DCAACCCCCCACCCAACAABBDBBDDDBDDDDDDBBDPCAACAAACAAAAAACDBBBBBBDBBBDBBDACCCAAABBBDDDAACACDBDBBCAACDBBDP", length: 4},
		{s: "AXFCCCCCACCACCCAACAACACACAACCAAABBBDDBBDBDBDBBDBBDDDBDDBDDDDDCCACDBDDCABD", length: 3},
		{s: "ACACACACACACACACACACACACACACACACACACACACACACACACACDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDB", length: 0},
	} {
		length := minLength(tc.s)
		assert.Equal(t, tc.length, length)
	}
}
