package greedy

/*
Key Observations:
- If the total number of cards is not divisible by groupSize,
	it's impossible to rearrange the cards into the desired groups.
- There may be duplicate cards present in the array,
	but a valid group does not contain duplicates.
- You don't need 2 groups if W=2. Only size of group needs to be 2.
	There can be any number of groups.
- Other main condition is, numbers in groups should be consecutive
	i.e. nextNumber = prevNumber + 1 in each group.
- For input = [2,1], W=2. We can have 1 group of size(W)=2
	where numbers are consecutive i.e [1,2]
- Consider following input - [8,10,12]. W=3. Here we can form 1 group of size=3,
	but numbers in it are not consecutive. Hence answer is false.
*/

// https://leetcode.com/problems/hand-of-straights/

// time: O(2n) = O(n)
// space: O(n)
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	var count = make(map[int]int)
	for _, card := range hand { // O(n)
		count[card]++
	}
	var start, curr int
	for _, card := range hand { // O(2n)
		// find the start of current group(s)
		start = card
		for count[start-1] > 0 { // O(n)
			start--
		}
		for start <= card { // O(n)
			for count[start] > 0 {
				for curr = start + groupSize - 1; curr >= start; curr-- {
					if count[curr] == 0 {
						return false
					}
					count[curr]--
				}
			}
			start++
		}
	}
	return true
}
