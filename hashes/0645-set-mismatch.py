import unittest
from typing import List

# https://leetcode.com/problems/set-mismatch/

# python3 -m unittest hashes/0645-set-mismatch.py


class Solution(unittest.TestCase):
    # # Approach 1: Hash Table (Brute Force)
    # # Intuition: Count frequency of each number. The duplicate appears twice,
    # # and the missing number doesn't appear at all.
    # # Time Complexity: O(n) - single pass to build frequency map, single pass to find results
    # # Space Complexity: O(n) - hash table stores up to n entries
    # def findErrorNums(self, nums: List[int]) -> List[int]:
    #     freq: Dict[int, int] = {}
    #     duplicate = missing = 0

    #     # Count frequencies
    #     for num in nums:
    #         freq[num] = freq.get(num, 0) + 1

    #     # Find duplicate and missing
    #     for i in range(1, len(nums) + 1):
    #         if i not in freq:
    #             missing = i
    #         elif freq[i] == 2:
    #             duplicate = i

    #     return [duplicate, missing]

    # # Approach 2: Mathematical (Optimized)
    # # Intuition: Use the mathematical properties:
    # # - Expected sum = n*(n+1)/2
    # # - Expected sum of squares = n*(n+1)*(2n+1)/6
    # # Let duplicate = d, missing = m
    # # - actual_sum - expected_sum = d - m ... (equation 1)
    # # - actual_sum_sq - expected_sum_sq = d^2 - m^2 = (d-m)(d+m) ... (equation 2)
    # # From eq2/eq1: d + m
    # # Solving: d = ((d-m) + (d+m))/2, m = ((d+m) - (d-m))/2
    # # Time Complexity: O(n) - single pass to calculate sums
    # # Space Complexity: O(1) - only using a few variables
    # def findErrorNums(self, nums: List[int]) -> List[int]:
    #     n = len(nums)
    #     expected_sum = n * (n + 1) // 2
    #     expected_sum_sq = n * (n + 1) * (2 * n + 1) // 6

    #     actual_sum = sum(nums)
    #     actual_sum_sq = sum(x * x for x in nums)

    #     diff_sum = actual_sum - expected_sum  # d - m
    #     diff_sum_sq = actual_sum_sq - expected_sum_sq  # d^2 - m^2

    #     sum_dm = diff_sum_sq // diff_sum  # d + m

    #     duplicate = (diff_sum + sum_dm) // 2
    #     missing = sum_dm - duplicate

    #     return [duplicate, missing]

    # Approach 3: In-place Marking (Most Optimal for Space)
    # Intuition: Use array indices to mark seen numbers by negating values.
    # When we see a number that's already marked (negative), it's the duplicate.
    # The index that remains positive indicates the missing number.
    # Time Complexity: O(n) - two passes through array
    # Space Complexity: O(1) - modifying input array in-place
    def findErrorNums(self, nums: List[int]) -> List[int]:
        duplicate = -1

        # Mark seen numbers by negating the value at their index
        for num in nums:
            index = abs(num) - 1  # Convert to 0-indexed
            if nums[index] < 0:
                # Already marked, this is the duplicate
                duplicate = abs(num)
            else:
                # Mark as seen
                nums[index] = -nums[index]

        # Find the missing number (index that's still positive)
        missing = -1
        for i in range(len(nums)):
            if nums[i] > 0:
                missing = i + 1
                break

        return [duplicate, missing]

    def testFindErrorNums(self) -> None:
        for nums, expected in [
            ([1, 2, 2, 4], [2, 3]),
            ([1, 1], [1, 2]),
            ([3, 2, 2], [2, 1]),
            ([2, 2], [2, 1]),
            ([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10], [10, 11]),
            ([1, 5, 3, 2, 2, 7, 6, 4, 8, 9], [2, 10]),
        ]:
            output = self.findErrorNums(nums.copy())
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
