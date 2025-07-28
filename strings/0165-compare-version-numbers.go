package strings

// https://leetcode.com/problems/compare-version-numbers/

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
