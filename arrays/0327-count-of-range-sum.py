from typing import List
import unittest

# https://leetcode.com/problems/count-of-range-sum/

# python3 -m unittest arrays/0327-count-of-range-sum.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # Generate all possible range sums and count those in [lower, upper]
    # # For each starting position i, compute all range sums starting from i
    # # Time: O(n^2) - nested loops to generate all ranges
    # # Space: O(n) - prefix sum array
    # def countRangeSum(self, nums: List[int], lower: int, upper: int) -> int:
    #     n = len(nums)
    #     # Build prefix sum array for O(1) range sum queries
    #     prefix = [0] * (n + 1)
    #     for i in range(n):
    #         prefix[i + 1] = prefix[i] + nums[i]
    
    #     count = 0
    #     # Try all possible ranges [i, j]
    #     for i in range(n):
    #         for j in range(i, n):
    #             # Range sum S(i, j) = prefix[j+1] - prefix[i]
    #             range_sum = prefix[j + 1] - prefix[i]
    #             if lower <= range_sum <= upper:
    #                 count += 1
    
    #     return count

    # Approach 2: Divide and Conquer with Merge Sort
    # Key insight: For range sum S(i,j) = prefix[j+1] - prefix[i]
    # We need: lower <= prefix[j+1] - prefix[i] <= upper
    # Rearranging: prefix[i] + lower <= prefix[j+1] <= prefix[i] + upper
    #
    # During merge sort of prefix sums:
    # - For each prefix[i] in left half, count how many prefix[j] in right half
    #   satisfy: prefix[i] + lower <= prefix[j] <= prefix[i] + upper
    # - Use two pointers since both halves are sorted
    #
    # Time: O(n log n) - merge sort with linear merge
    # Space: O(n) - for prefix array and merge temporary array
    def countRangeSum(self, nums: List[int], lower: int, upper: int) -> int:
        # Build prefix sum array
        n = len(nums)
        prefix = [0] * (n + 1)
        for i in range(n):
            prefix[i + 1] = prefix[i] + nums[i]

        def merge_sort(start: int, end: int) -> int:
            if end - start <= 1:
                return 0

            mid = (start + end) // 2
            count = merge_sort(start, mid) + merge_sort(mid, end)

            # Count valid pairs where i in [start, mid) and j in [mid, end)
            # We need: lower <= prefix[j] - prefix[i] <= upper
            # Which means: prefix[i] + lower <= prefix[j] <= prefix[i] + upper
            j_low = j_high = mid
            for i in range(start, mid):
                # Find the range [j_low, j_high) where prefix[j] satisfies the condition
                while j_low < end and prefix[j_low] - prefix[i] < lower:
                    j_low += 1
                while j_high < end and prefix[j_high] - prefix[i] <= upper:
                    j_high += 1
                count += j_high - j_low

            # Merge the two sorted halves
            sorted_arr: List[int] = []
            left, right = start, mid
            while left < mid and right < end:
                if prefix[left] <= prefix[right]:
                    sorted_arr.append(prefix[left])
                    left += 1
                else:
                    sorted_arr.append(prefix[right])
                    right += 1
            sorted_arr.extend(prefix[left:mid])
            sorted_arr.extend(prefix[right:end])

            # Copy back to prefix array
            prefix[start:end] = sorted_arr

            return count

        return merge_sort(0, n + 1)

    def test(self):
        for nums, lower, upper, expected in [
            ([-2, 5, -1], -2, 2, 3),
            ([0], 0, 0, 1),
            ([2147483647, -2147483648, -1, 0], -1, 0, 4),
            ([0, -3, -3, 1, 1, 2], 3, 5, 2),
            ([-2, 5, -1], 0, 2, 1),
        ]:
            output = self.countRangeSum(nums[:], lower, upper)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
