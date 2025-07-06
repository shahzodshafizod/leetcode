from typing import List
import unittest

# https://leetcode.com/problems/count-subarrays-with-score-less-than-k/


class Solution(unittest.TestCase):
    # Approach: Sliding Window
    # Time: O(n)
    # Space: O(1)
    def countSubarrays(self, nums: List[int], k: int) -> int:
        count, total, start = 0, 0, 0
        for end, num in enumerate(nums):
            total += num
            while total * (end - start + 1) >= k:
                total -= nums[start]
                start += 1
            count += end - start + 1
        return count

    def test(self):
        for nums, k, expected in [
            ([2, 1, 4, 3, 5], 10, 6),
            ([1, 1, 1], 5, 5),
        ]:
            output = self.countSubarrays(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
