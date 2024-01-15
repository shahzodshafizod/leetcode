package stacks

/*
Problem:
Given a string only containing round blackets '(' and ')' and lowercase characters,
remove the least amount of brackets so the string is valid.
A string is considered valid if it is empty or if there are brackets, they all close.

Step 1: Verify the constraints
	- What do we return from our algorithm?
		: Return a valid string with the fewest brackets removed.
	- Will there be spaces in the string?
		: No, the string only contains lowercase characters, '(' and ')'.
	- Is a string containing only lowercase characters valid?
		: Yes, you don't need any brackets for a string to be valid.
Step 2: Write out some test cases
	- "a)bc(d)" -> "abc(d)"
	- "(ab(c)d" -> "ab(c)d" | "(abc)d"
	- "))((" -> ""
*/

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

func minRemoveToMakeValid(s string) string {
	var slice = []byte(s)
	var remove = make([]int, 0, len(slice))
	for idx, elem := range slice {
		switch elem {
		case '(':
			remove = append(remove, idx)
		case ')':
			if len := len(remove); len > 0 && slice[remove[len-1]] == '(' {
				remove = remove[:len-1]
			} else {
				remove = append(remove, idx)
			}
		}
	}
	for idx := len(remove) - 1; idx >= 0; idx-- {
		slice = append(slice[:remove[idx]], slice[remove[idx]+1:]...)
	}
	return string(slice)
}
