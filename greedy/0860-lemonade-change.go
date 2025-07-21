package greedy

// https://leetcode.com/problems/lemonade-change/

func lemonadeChange(bills []int) bool {
	fiveStack, tenStack := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			fiveStack++
		case 10:
			tenStack++
			fiveStack--
		case 20:
			if tenStack > 0 {
				tenStack--
				fiveStack--
			} else {
				fiveStack -= 3
			}
		}
		if fiveStack < 0 {
			return false
		}
	}
	return true
}
