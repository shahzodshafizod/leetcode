package stacks

// https://leetcode.com/problems/crawler-log-folder/

func minOperations(logs []string) int {
	stack := 0
	for _, log := range logs {
		switch log {
		case "../":
			stack = max(stack-1, 0)
		case "./":
		default:
			stack++
		}
	}
	return stack
}
