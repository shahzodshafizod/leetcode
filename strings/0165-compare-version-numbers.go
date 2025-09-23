package strings

// https://leetcode.com/problems/compare-version-numbers/

// func compareVersion(version1 string, version2 string) int {
// 	revs1 := strings.Split(version1, ".")
// 	revs2 := strings.Split(version2, ".")
// 	n1, n2 := len(revs1), len(revs2)

// 	var rev1, rev2 int

// 	for i1, i2 := 0, 0; i1 < n1 || i2 < n2; i1, i2 = i1+1, i2+1 {
// 		rev1, rev2 = 0, 0
// 		if i1 < n1 {
// 			rev1, _ = strconv.Atoi(revs1[i1])
// 		}

// 		if i2 < n2 {
// 			rev2, _ = strconv.Atoi(revs2[i2])
// 		}

// 		if rev1 < rev2 {
// 			return -1
// 		} else if rev1 > rev2 {
// 			return 1
// 		}
// 	}

// 	return 0
// }

// Approach: Two-Pointers
// time: O(n1 + n2)
// space: O(1)
func compareVersion(version1 string, version2 string) int {
	idx1, idx2 := 0, 0
	n1, n2 := len(version1), len(version2)

	var revision1, revision2 int

	for idx1 < n1 || idx2 < n2 {
		revision1, revision2 = 0, 0
		for idx1 < n1 && version1[idx1] != '.' {
			revision1 = revision1*10 + int(version1[idx1]-'0')
			idx1++
		}

		for idx2 < n2 && version2[idx2] != '.' {
			revision2 = revision2*10 + int(version2[idx2]-'0')
			idx2++
		}

		if revision1 < revision2 {
			return -1
		} else if revision1 > revision2 {
			return 1
		}

		idx1++
		idx2++
	}

	return 0
}
