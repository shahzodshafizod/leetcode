package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestRemoveOccurrences$
func TestRemoveOccurrences(t *testing.T) {
	for _, tc := range []struct {
		s       string
		part    string
		cleared string
	}{
		{s: "a", part: "aa", cleared: "a"},
		{s: "eemckxmckx", part: "emckx", cleared: ""},
		{s: "ccctltctlltlb", part: "ctl", cleared: "b"},
		{s: "yjyjqnaxlbqnaxlbfss", part: "yjqnaxlb", cleared: "fss"},
		{s: "gjzgbpggjzgbpgsvpwdk", part: "gjzgbpg", cleared: "svpwdk"},
		{s: "hhvhvaahvahvhvaavhvaasshvahvaln", part: "hva", cleared: "ssln"},
		{s: "daabcbaabcbc", part: "abc", cleared: "dab"},
		{s: "axxxxyyyyb", part: "xy", cleared: "ab"},
		{s: "wwwwwwwwwwwwwwwwwwwwwvwwwwswxwwwwsdwxweeohapwwzwuwajrnogb", part: "w", cleared: "vsxsdxeeohapzuajrnogb"},
		// {s: "kpygkivtlqoockpygkivtlqoocssnextkqzjpycbylkaondsskpygkpygkivtlqoocssnextkqzjpkpygkivtlqoocssnextkqzjpycbylkaondsycbylkaondskivtlqoocssnextkqzjpycbylkaondssnextkqzjpycbylkaondshijzgaovndkjiiuwjtcpdpbkrfsi", part: "kpygkivtlqoocssnextkqzjpycbylkaonds", cleared: "hijzgaovndkjiiuwjtcpdpbkrfsi"},
	} {
		cleared := removeOccurrences(tc.s, tc.part)
		assert.Equal(t, tc.cleared, cleared)
	}
}
