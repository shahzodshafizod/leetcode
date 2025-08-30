package hashes

// https://leetcode.com/problems/finding-3-digit-even-numbers/

func findEvenNumbers(digits []int) []int {
	var count [10]int
	for _, digit := range digits {
		count[digit]++
	}

	var uniques []int

	for sadi := 1; sadi < 10; sadi++ {
		if count[sadi] == 0 {
			continue
		}

		count[sadi]--
		for dahi := range 10 {
			if count[dahi] == 0 {
				continue
			}

			count[dahi]--
			for vohid := 0; vohid < 10; vohid += 2 {
				if count[vohid] == 0 {
					continue
				}

				count[vohid]--
				uniques = append(uniques, sadi*100+dahi*10+vohid)
				count[vohid]++
			}

			count[dahi]++
		}

		count[sadi]++
	}

	return uniques
}
