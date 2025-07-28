package arrays

// https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	for idx, length := 0, len(flowerbed); idx < length && n > 0; idx++ {
		if (idx == 0 || flowerbed[idx-1] == 0) &&
			flowerbed[idx] == 0 &&
			(idx == length-1 || flowerbed[idx+1] == 0) {
			flowerbed[idx] = 1
			n--
		}
	}

	return n == 0
}
