package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMinWindow$
func TestMinWindow(t *testing.T) {
	for _, tc := range []struct {
		s      string
		t      string
		window string
	}{
		{s: "ADOBECODEBANC", t: "ABC", window: "BANC"},
		{s: "a", t: "a", window: "a"},
		{s: "a", t: "aa", window: ""},
		{s: "aa", t: "a", window: "a"},
		{s: "aaaaaaaaaaaabbbbbcdd", t: "abcdd", window: "abbbbbcdd"},
		{s: "caccaacaaaabbcaccaccc", t: "acccacbccc", window: "caaaabbcaccaccc"},
		{s: "cbbbacccccbbbacbabbabacbabbbabaacbaccccbcbcbcca", t: "abcbcabaacccababacbabcacbc", window: "acbabbabacbabbbabaacbaccccbcbc"},
		{s: "babb", t: "baba", window: ""},
		{s: "A", t: "", window: ""},
		{s: "wegdtzwabazduwwdysdetrrctotpcepalxdewzezbfewbabbseinxbqqplitpxtcwwhuyntbtzxwzyaufihclztckdwccpeyonumbpnuonsnnsjscrvpsqsftohvfnvtbphcgxyumqjzltspmphefzjypsvugqqjhzlnylhkdqmolggxvneaopadivzqnpzurmhpxqcaiqruwztroxtcnvhxqgndyozpcigzykbiaucyvwrjvknifufxducbkbsmlanllpunlyohwfsssiazeixhebipfcdqdrcqiwftutcrbxjthlulvttcvdtaiwqlnsdvqkrngvghupcbcwnaqiclnvnvtfihylcqwvderjllannflchdklqxidvbjdijrnbpkftbqgpttcagghkqucpcgmfrqqajdbynitrbzgwukyaqhmibpzfxmkoeaqnftnvegohfudbgbbyiqglhhqevcszdkokdbhjjvqqrvrxyvvgldtuljygmsircydhalrlgjeyfvxdstmfyhzjrxsfpcytabdcmwqvhuvmpssingpmnpvgmpletjzunewbamwiirwymqizwxlmojsbaehupiocnmenbcxjwujimthjtvvhenkettylcoppdveeycpuybekulvpgqzmgjrbdrmficwlxarxegrejvrejmvrfuenexojqdqyfmjeoacvjvzsrqycfuvmozzuypfpsvnzjxeazgvibubunzyuvugmvhguyojrlysvxwxxesfioiebidxdzfpumyon", t: "ozgzyywxvtublcl", window: "tcnvhxqgndyozpcigzykbiaucyvwrjvknifufxducbkbsmlanl"},
		{s: "cabwefgewcwaefgcf", t: "cae", window: "cwae"},
	} {
		window := minWindow(tc.s, tc.t)
		assert.Equal(t, tc.window, window)
	}
}
