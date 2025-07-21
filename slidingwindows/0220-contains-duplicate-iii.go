package slidingwindows

import "math"

/*
Bucket Sort:

First, we distribute each element of an array into a number of buckets.
In our case, the bucket looks like [0,t],[t+1,2t+1],….

Each time we distribute the current element nums[i] into a bucket,
if this bucket is not empty then we find two elements
that satisfy the two above requirements and we are done;
otherwise we check two adjacent buckets to see whether any one of them is non empty
and it contains an element whose value lies in the range of [nums[i]-t,nums[i]+t].

Some Intricacies:
- Since the problem defines the difference between two elements as absolute,
	the difference is never to be negatives
- Also, rather than use interval [0,t], we redefine the equivalent interval
	to be [0,t+1). In this way when t is equal to zero,
	we do not have to write extra code to take care of this special case.
- Third, we use type long when computing the absolute difference here to avoid overflows.
	Consider this test case: nums[i]=-1 while nums[j] =Integer.MAX_VALUE.
*/

// https://leetcode.com/problems/contains-duplicate-iii/

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) <= 1 || valueDiff < 0 {
		return false
	}
	width := int64(valueDiff) + 1
	getbucketkey := func(num int64) int64 {
		if num < 0 {
			return (num+1)/width - 1
		}
		return num / width
	}
	buckets := make(map[int64]int64)
	var bucketkey, num int64
	for idx := range nums {
		num = int64(nums[idx])
		bucketkey = getbucketkey(num)
		if _, exists := buckets[bucketkey]; exists {
			return true
		}
		if _, exists := buckets[bucketkey-1]; exists &&
			math.Abs(float64(num-buckets[bucketkey-1])) < float64(width) {
			return true
		}
		if _, exists := buckets[bucketkey+1]; exists &&
			math.Abs(float64(num-buckets[bucketkey+1])) < float64(width) {
			return true
		}
		buckets[bucketkey] = num
		if idx-indexDiff >= 0 {
			delete(buckets, getbucketkey(int64(nums[idx-indexDiff])))
		}
	}
	return false
}

// // time: O(n^2)
// // space: O(1)
// func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
// 	var n = len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n && j <= i+indexDiff; j++ {
// 			if math.Abs(float64(nums[i]-nums[j])) <= float64(valueDiff) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
