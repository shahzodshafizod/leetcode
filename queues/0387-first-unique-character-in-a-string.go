package queues

// https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	var unique [26][2]int
	n := len(s)
	for i, len := 0, n; i < len; i++ {
		unique[s[i]-'a'] = [2]int{unique[s[i]-'a'][0] + 1, i}
	}
	index := n
	for _, item := range unique {
		if item[0] == 1 && item[1] < index {
			index = item[1]
		}
	}
	if index == n {
		index = -1
	}
	return index
}

// func firstUniqChar(s string) int {
// 	var unique [26]int
// 	var n = len(s)
// 	for i := 0; i < n; i++ {
// 		unique[s[i]-'a']++
// 	}
// 	for i := 0; i < n; i++ {
// 		if unique[s[i]-'a'] == 1 {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func firstUniqChar(s string) int {
// 	var unique = make(map[rune]int)
// 	for _, r := range s {
// 		unique[r] = unique[r] + 1
// 	}
// 	for index, r := range s {
// 		if unique[r] == 1 {
// 			return index
// 		}
// 	}
// 	return -1
// }
