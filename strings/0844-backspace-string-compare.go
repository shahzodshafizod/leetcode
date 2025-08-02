package strings

/*
Problem:
Given two strings S and T, return if they equal when both are typed out.
Any '#' that appears in the string counts as a backspace

Step 1: Verify the constraints
	- What happens when two #'s appear beside each other?
		: Delete the two values before the first #.
	- What happens to # when there is no character to remove?
		: It deletes nothing then, just like backspace would.
	- Are two empty strings equal to each other?
		: Yes, consider two empty strings as equal.
	- Does case sensitivity matter?
		: Yes it does, "a" does not equal "A".
Step 2: Write out some test cases
	- "ab#z", "az#z" -> true
	- "abc#d", "acc#d" -> false
	- "x#y#z#", "a#" -> true
	- "a###b", "b" -> true
	- "Ab#z", "ab#z" -> false
Step 8: Can we optimize our solution?
	- 1st Hint: Utilize the original strings.
	- 2nd Hint: Use the 2 pointer technique.
	- 3rd Hint: Start from the end of the strings!
*/

// https://leetcode.com/problems/backspace-string-compare/

func backspaceCompare(s string, t string) bool {
	equal := true

	var skips int

	for indexS, indexT := len(s)-1, len(t)-1; ; {
		for indexS >= 0 && s[indexS] == '#' {
			skips = 0

			for indexS >= 0 && s[indexS] == '#' {
				indexS--
				skips++
			}

			for skips > 0 && indexS >= 0 {
				if s[indexS] == '#' {
					skips += 2
				}

				indexS--
				skips--
			}
		}

		for indexT >= 0 && t[indexT] == '#' {
			skips = 0

			for indexT >= 0 && t[indexT] == '#' {
				indexT--
				skips++
			}

			for skips > 0 && indexT >= 0 {
				if t[indexT] == '#' {
					skips += 2
				}

				indexT--
				skips--
			}
		}

		if indexS < 0 && indexT < 0 {
			break
		}

		if indexS < 0 || indexT < 0 {
			equal = false

			break
		}

		if s[indexS] != t[indexT] {
			equal = false

			break
		}

		indexS--
		indexT--
	}

	return equal
}

// func backspaceCompare(s string, t string) bool {
// 	runeS := removeHashs(s)
// 	runeT := removeHashs(t)
// 	if len(runeS) != len(runeT) {
// 		return false
// 	}
// 	for i := len(runeS) - 1; i >= 0; i-- {
// 		if runeS[i] != runeT[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func removeHashs(str string) []rune {
// 	var result = make([]rune, 0)
// 	_ = strings.Map(func(r rune) rune {
// 		if r != '#' {
// 			result = append(result, r)
// 		} else {
// 			if length := len(result); length > 0 {
// 				result = result[:length-1]
// 			}
// 		}
// 		return r
// 	}, str)
// 	return result
// }
