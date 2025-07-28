package strings

import "strings"

// https://leetcode.com/problems/text-justification/

// Approach: Brute-Force
// Time: O(N*maxWidth)
// Space: O(maxWidth)
func fullJustify(words []string, maxWidth int) []string {
	lines := make([]string, 0)
	line, width := make([]string, 0), 0

	var wcount, spaces, extra int

	for idx, n := 0, len(words); idx < n; idx++ {
		if width+len(line)+len(words[idx]) > maxWidth {
			// justify
			wcount = len(line)
			spaces = maxWidth - width
			extra = spaces % max(1, wcount-1)

			spaces /= max(1, wcount-1)
			for j := 0; j < wcount-1; j++ {
				line[j] += strings.Repeat(" ", spaces)
				if extra > 0 {
					line[j] += " "
					extra--
				}
			}

			if wcount == 1 {
				line[0] += strings.Repeat(" ", spaces)
			}

			lines = append(lines, strings.Join(line, ""))
			line, width = []string{}, 0 // reset
		}

		line = append(line, words[idx])
		width += len(words[idx])
	}

	lastLine := strings.Join(line, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	lines = append(lines, lastLine)

	return lines
}
