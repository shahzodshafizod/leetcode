package maths

// https://leetcode.com/problems/count-symmetric-integers/

func countSymmetricIntegers(low int, high int) int {
	// calculate length of low
	// and the next length's start number
	var length, next = 0, 1
	for tmp := low; tmp > 0; tmp /= 10 {
		length++
		next *= 10
	}
	var sum, tmp int
	var count = 0
	for num := low; num <= high; num++ {
		// if number became one digit longer
		// we change length and the next length's start number
		if num == next {
			length++
			next *= 10
		}
		// if length is odd, it cannot be split into two parts
		// so we can skip the whole current odd-length section
		if length&1 == 1 {
			num = next - 1
			continue
		}
		// now, we add right part digits to total
		sum, tmp = 0, num
		for idx := length / 2; idx > 0; idx-- {
			sum += tmp % 10
			tmp /= 10
		}
		// but substract left digits from total
		for ; tmp > 0; tmp /= 10 {
			sum -= tmp % 10
		}
		// if left and right parts have equal sums, total should be 0
		if sum == 0 {
			count++
		}
	}
	return count
}
