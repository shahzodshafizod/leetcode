from typing import List
import unittest

# https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/

# python3 -m unittest twopointers/1498-number-of-subsequences-that-satisfy-the-given-sum-condition.py

class Solution(unittest.TestCase):
    # Approach: Two Pointer
    # Time: O(nlogn)
    # Space: O(1)
    def numSubseq(self, nums: List[int], target: int) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        power2 = [0] * n
        power2[0] = 1
        for idx in range(1, n):
            power2[idx] = (power2[idx-1] * 2) % MOD
        nums.sort()
        left, right = 0, n-1
        count = 0
        while left <= right:
            if nums[left]+nums[right] > target:
                right -= 1
            else:
                # for each element in the subarray [left+1:right]
                # we can pick or not pick, so there are
                # 2 ^ (right - left) subsequences in total
                count = (count + power2[right-left]) % MOD
                left += 1
        return count

    def test(self):
        for nums, target, expected in [
            ([3,5,6,7], 9, 4),
            ([3,3,6,8], 10, 6),
            ([2,3,3,4,6,7], 12, 61),
            ([7,10,7,3,7,5,4], 12, 56),
        ]:
            output = self.numSubseq(nums, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
