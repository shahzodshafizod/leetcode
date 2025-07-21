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
	tortoise, hare := 0, 0
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

// // Bit Manipulation
// func findDuplicate(nums []int) int {
// 	var duplicate, number, count int
// 	for _, pointer := range nums {
// 		pointer--
// 		pointer &= 0xFFFF
// 		number = nums[pointer] & 0xFFFF
// 		count = nums[pointer] >> 16
// 		count++
// 		if count == 2 {
// 			duplicate = pointer + 1
// 			break
// 		}
// 		nums[pointer] = (count << 16) + number
// 	}
// 	for idx := range nums {
// 		nums[idx] = ((nums[idx] - 1) & 0xFFFF) + 1
// 	}
// 	return duplicate
// }

// func findDuplicate(nums []int) int {
// 	var duplicate int
// 	for _, num := range nums {
// 		if num < 0 {
// 			num = -num
// 		}
// 		if nums[num] < 0 {
// 			duplicate = num
// 			break
// 		}
// 		nums[num] = -nums[num]
// 	}
// 	for idx := range nums {
// 		if nums[idx] < 0 {
// 			nums[idx] = -nums[idx]
// 		}
// 	}
// 	return duplicate
// }
