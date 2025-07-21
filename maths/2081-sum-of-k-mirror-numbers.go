package maths

// https://leetcode.com/problems/sum-of-k-mirror-numbers/

// Time: O(sqrt(10))
// Space: O(1)
func kMirror(k int, n int) int64 {
	var digits [100]int64
	isBaseKMirror := func(num int64) bool {
		right := -1
		for ; num > 0; num /= int64(k) {
			right++
			digits[right] = num % int64(k)
		}
		for left := 0; left < right; {
			if digits[left] != digits[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	var left, right, i int
	var mirror, tmp int64
	var sum int64 = 0
	for left = 1; n > 0; left = right {
		right = left * 10
		for p := 0; p < 2; p++ {
			for i = left; i < right && n > 0; i++ {
				mirror, tmp = int64(i), int64(i)
				if p == 0 {
					tmp /= 10 // first odd-length, then even-length
				}
				for ; tmp > 0; tmp /= 10 {
					mirror = mirror*10 + tmp%10
				}
				if isBaseKMirror(mirror) {
					sum += mirror
					n--
				}
			}
		}
	}
	return sum
}
