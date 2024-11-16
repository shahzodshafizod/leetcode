package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumWays1639$
func TestNumWays1639(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		target string
		ways   int
	}{
		{words: []string{"acca", "bbbb", "caca"}, target: "aba", ways: 6},
		{words: []string{"abba", "baab"}, target: "bab", ways: 4},
		// {words: []string{"cbabddddbc", "addbaacbbd", "cccbacdccd", "cdcaccacac", "dddbacabbd", "bdbdadbccb", "ddadbacddd", "bbccdddadd", "dcabaccbbd", "ddddcddadc", "bdcaaaabdd", "adacdcdcdd", "cbaaadbdbb", "bccbabcbab", "accbdccadd", "dcccaaddbc", "cccccacabd", "acacdbcbbc", "dbbdbaccca", "bdbddbddda", "daabadbacb", "baccdbaada", "ccbabaabcb", "dcaabccbbb", "bcadddaacc", "acddbbdccb", "adbddbadab", "dbbcdcbcdd", "ddbabbadbb", "bccbcbbbab", "dabbbdbbcb", "dacdabadbb", "addcbbabab", "bcbbccadda", "abbcacadac", "ccdadcaada", "bcacdbccdb"}, target: "bcbbcccc", ways: 677452090},
	} {
		ways := numWays1639(tc.words, tc.target)
		assert.Equal(t, tc.ways, ways)
	}
}
