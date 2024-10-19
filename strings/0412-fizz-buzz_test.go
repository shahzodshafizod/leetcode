package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestFizzBuzz$
func TestFizzBuzz(t *testing.T) {
	for _, tc := range []struct {
		n      int
		answer []string
	}{
		{n: 3, answer: []string{"1", "2", "Fizz"}},
		{n: 5, answer: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{n: 15, answer: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	} {
		answer := fizzBuzz(tc.n)
		assert.Equal(t, tc.answer, answer)
	}
}
