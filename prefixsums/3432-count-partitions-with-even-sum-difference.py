from typing import List
import unittest

# https://leetcode.com/problems/count-partitions-with-even-sum-difference/

# python3 -m unittest prefixsums/3432-count-partitions-with-even-sum-difference.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force - Calculate Sums for Each Partition
    # # For each partition, recalculate left and right sums
    # # Time Complexity: O(n²)
    # # Space Complexity: O(1)
    # def countPartitions(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     count = 0

    #     for i in range(n - 1):
    #         left_sum = sum(nums[: i + 1])
    #         right_sum = sum(nums[i + 1 :])

    #         if (left_sum - right_sum) % 2 == 0:
    #             count += 1

    #     return count

    # Approach 2: Prefix Sum (Optimal)
    # Use running sums to avoid recalculation
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def countPartitions(self, nums: List[int]) -> int:
        n = len(nums)
        if n < 2:
            return 0

        # Calculate total sum
        right_sum = sum(nums)
        left_sum = 0
        count = 0

        # Try each partition point
        for i in range(n - 1):
            left_sum += nums[i]
            right_sum -= nums[i]

            # Check if difference is even
            if (left_sum - right_sum) % 2 == 0:
                count += 1

        return count

    def test(self):
        for nums, expected in [
            ([1, 2, 2], 0),
            ([2, 4, 6, 8], 3),
            ([1, 1, 1, 1], 3),
            ([1, 2, 3], 2),
            ([10, 20], 1),
        ]:
            output = self.countPartitions(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
