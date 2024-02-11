package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestNumberToWords$
func TestNumberToWords(t *testing.T) {
	for _, tc := range []struct {
		num            int
		representation string
	}{
		{num: 123, representation: "One Hundred Twenty Three"},
		{num: 12345, representation: "Twelve Thousand Three Hundred Forty Five"},
		{num: 412345, representation: "Four Hundred Twelve Thousand Three Hundred Forty Five"},
		{num: 1234567, representation: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{num: 1000, representation: "One Thousand"},
		{num: 100001, representation: "One Hundred Thousand One"},
		{num: 100, representation: "One Hundred"},
		{num: 2147483647, representation: "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven"},
		{num: 8342147483647, representation: "Eight Trillion Three Hundred Forty Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven"},
		{num: 30, representation: "Thirty"},
		{num: 130, representation: "One Hundred Thirty"},
		{num: 0, representation: "Zero"},
		{num: 1, representation: "One"},
		{num: 222, representation: "Two Hundred Twenty Two"},
		{num: 1234, representation: "One Thousand Two Hundred Thirty Four"},
		{num: 31337, representation: "Thirty One Thousand Three Hundred Thirty Seven"},
		{num: 100100, representation: "One Hundred Thousand One Hundred"},
		{num: 200111, representation: "Two Hundred Thousand One Hundred Eleven"},
		{num: 1234567891, representation: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"},
		{num: 151, representation: "One Hundred Fifty One"},
		{num: 1414312, representation: "One Million Four Hundred Fourteen Thousand Three Hundred Twelve"},
		{num: 1241234113, representation: "One Billion Two Hundred Forty One Million Two Hundred Thirty Four Thousand One Hundred Thirteen"},
	} {
		representation := numberToWords(tc.num)
		assert.Equal(t, tc.representation, representation)
	}
}

// go test -v -count=1 ./strings/ -run ^TestNumberToWordsReverse$
func TestNumberToWordsReverse(t *testing.T) {
	for _, tc := range []struct {
		num            int
		representation string
	}{
		{num: 123, representation: "Three Hundred Twenty One"},
		{num: 12345, representation: "Fifty Four Thousand Three Hundred Twenty One"},
		{num: 123456, representation: "Six Hundred Fifty Four Thousand Three Hundred Twenty One"},
		{num: 1234567, representation: "Seven Million Six Hundred Fifty Four Thousand Three Hundred Twenty One"},
		{num: 1000, representation: "One"},
		{num: 100001, representation: "One Hundred Thousand One"},
		{num: 100004, representation: "Four Hundred Thousand One"},
		{num: 100004002, representation: "Two Hundred Million Four Hundred Thousand One"}, // -
		{num: 10000422, representation: "Twenty Two Million Four Hundred Thousand One"},   // -
		{num: 0, representation: "Zero"},
	} {
		representation := numberToWordsReverse(tc.num)
		assert.Equal(t, tc.representation, representation)
	}
}
