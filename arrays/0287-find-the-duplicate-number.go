package arrays

/*
Follow up:
	- How can we prove that at least one duplicate number must exist in nums?
	- Can you solve the problem in linear runtime complexity?

Hints:
The follow-up questions hint that:
	1. It's a linked list problem
	2. We have to find, if a linked list has a cycle
	3. Floyd's Tortoise and Hare (Cycle Detection)

If all numbers from 1 to n existed, then
(total sum of nums' numbers) - (n*(n+1)/2)
could give us that duplicate
*/

// https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	// 1. find meeting point
	var tortoise, hare = 0, 0
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}
	// 2. find the duplicate
	hare = 0 // here, hare is tortoise:)
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}
	return tortoise
}
