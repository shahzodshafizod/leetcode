package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCrackSafe$
func TestCrackSafe(t *testing.T) {
	verifySolution := func(s string, n int, k int) bool {
		// Helper: power function
		pow := func(base, exp int) int {
			result := 1
			for range exp {
				result *= base
			}

			return result
		}

		// Check expected length
		totalPasswords := pow(k, n)

		expectedLength := totalPasswords + n - 1
		if len(s) != expectedLength {
			return false
		}

		// Generate all possible n-digit passwords
		passwords := make(map[string]bool)

		for i := range totalPasswords {
			password := ""

			num := i
			for range n {
				password = string(byte('0'+(num%k))) + password
				num /= k
			}

			passwords[password] = true
		}

		// Check if all passwords appear in result
		found := make(map[string]bool)
		for i := 0; i <= len(s)-n; i++ {
			found[s[i:i+n]] = true
		}

		// Verify all passwords are found
		return len(found) == len(passwords)
	}

	for _, tc := range []struct {
		n      int
		k      int
		verify func(string, int, int) bool
	}{
		{n: 1, k: 2, verify: verifySolution},
		{n: 2, k: 2, verify: verifySolution},
		{n: 2, k: 3, verify: verifySolution},
		{n: 3, k: 2, verify: verifySolution},
		{n: 2, k: 4, verify: verifySolution},
		{n: 3, k: 3, verify: verifySolution},
		{n: 4, k: 2, verify: verifySolution},
	} {
		result := crackSafe(tc.n, tc.k)
		assert.True(t, tc.verify(result, tc.n, tc.k))
	}
}
