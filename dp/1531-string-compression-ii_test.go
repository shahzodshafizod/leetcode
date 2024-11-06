package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestGetLengthOfOptimalCompression$
func TestGetLengthOfOptimalCompression(t *testing.T) {
	for _, tc := range []struct {
		s      string
		k      int
		length int
	}{
		{s: "x", k: 1, length: 0},
		{s: "aaabcccd", k: 2, length: 4},
		{s: "aabbaa", k: 2, length: 2},
		{s: "aaaaaaaaaaa", k: 0, length: 3},
		{s: "llllllllllttttttttt", k: 1, length: 4},
		{s: "eoongjjkjfelnkgkjohfjfjfhkmnmmlinkihhlfipgoejiniol", k: 13, length: 32},
		{s: "kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk", k: 41, length: 2},
		{s: "ohoohhoooohhohoooohhhohoohhoohooohhhoohhhooohohhooohhoohhhhhhhooooohhoooohooohooohhohhhhhhohohoohhoo", k: 74, length: 3},
		{s: "earmaypeacebeonearthmerrychristmashohohoooohappyhanukkahhappyholidayshappyholidaysandahappynewyearma", k: 20, length: 70},
		{s: "mkkoqqilqminmrgqprnokiknnrohlqggqniopnpgonjojqihrihqkkhikjgmmggolgmnhoinqqnmlqkoqoihpkkhppoiligjrjnm", k: 64, length: 24},
		{s: "llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll", k: 25, length: 3},
	} {
		length := getLengthOfOptimalCompression(tc.s, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
