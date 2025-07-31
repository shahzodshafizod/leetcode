package arrays

// https://leetcode.com/problems/statistics-from-a-large-sample/

func sampleStats(count []int) []float64 {
	mi, ma := -1, 256

	var sum float64

	mode := [2]int{0, 0}

	var leftCount, rightCount uint64

	var mid int

	left, right := 0, 255
	for left <= right {
		if count[left] == 0 {
			left++
			continue
		}

		if count[right] == 0 {
			right--
			continue
		}

		if leftCount < rightCount {
			sampleStatsSet(left, count[left], &sum, &mi, &ma, &mode)
			leftCount += uint64(count[left])
			mid = left
			left++
		} else {
			sampleStatsSet(right, count[right], &sum, &mi, &ma, &mode)
			rightCount += uint64(count[right])
			mid = right
			right--
		}
	}

	total := leftCount + rightCount
	half := total / 2

	var neighbor int

	if leftCount == rightCount {
		if mid == left {
			neighbor = sampleStatsNextRight(count, right)
		} else {
			neighbor = sampleStatsNextLeft(count, left)
		}
	} else {
		neighbor = mid

		if leftCount <= half {
			mid = sampleStatsNextLeft(count, left)
		} else {
			mid = sampleStatsNextRight(count, right)
		}

		if count[mid] > 1 {
			neighbor = mid
		}
	}

	med := float64(mid)
	if total%2 == 0 {
		med = float64(mid+neighbor) / 2
	}

	mea := sum / float64(total)
	mo := float64(mode[1])

	return []float64{float64(mi), float64(ma), mea, med, mo}
}

func sampleStatsSet(num int, count int, sum *float64, minimum, maximum *int, mode *[2]int) {
	*sum += float64(num) * float64(count)

	if *minimum == -1 || num < *minimum {
		*minimum = num
	}

	if *maximum == 256 || num > *maximum {
		*maximum = num
	}

	if count > (*mode)[0] {
		*mode = [2]int{count, num}
	}
}

func sampleStatsNextLeft(count []int, left int) int {
	for left < 256 && count[left] == 0 {
		left++
	}

	return left
}

func sampleStatsNextRight(count []int, right int) int {
	for right >= 0 && count[right] == 0 {
		right--
	}

	return right
}
