import unittest
from typing import List
from bisect import bisect_left
from itertools import combinations

# https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/

# python3 -m unittest dp/2035-partition-array-into-two-arrays-to-minimize-sum-difference.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Try all 2^(2n) partitions)
    # # Time: O(2^(2n)) - exponential, trying all possible partitions
    # # Space: O(2n) - recursion stack
    # def minimumDifference(self, nums: List[int]) -> int:
    #     n = len(nums) // 2
    #     total = sum(nums)
    #     min_diff = 1 << 31
    
    #     def backtrack(idx: int, count: int, current_sum: int):
    #         nonlocal min_diff
    
    #         # We need exactly n elements in first partition
    #         if count > n or (len(nums) - idx) < (n - count):
    #             return
    
    #         if idx == len(nums):
    #             if count == n:
    #                 # Sum of second partition
    #                 other_sum = total - current_sum
    #                 diff = abs(current_sum - other_sum)
    #                 min_diff = min(min_diff, diff)
    #             return
    
    #         # Include current element in first partition
    #         backtrack(idx + 1, count + 1, current_sum + nums[idx])
    
    #         # Exclude current element (goes to second partition)
    #         backtrack(idx + 1, count, current_sum)
    
    #     backtrack(0, 0, 0)
    #     return min_diff

    # Approach #2: Meet in the Middle + Binary Search
    # Time: O(2^n * n) - generate all subsets for each half, then binary search
    # Space: O(2^n) - storing all possible sums for each half
    def minimumDifference(self, nums: List[int]) -> int:
        n = len(nums) // 2
        total = sum(nums)
        target = total // 2

        # Split array into two halves
        left_half = nums[:n]
        right_half = nums[n:]

        # For each half, generate all possible sums grouped by count of elements
        # left[k] = list of all possible sums using exactly k elements from left half
        left_sums: List[List[int]] = [[] for _ in range(n + 1)]
        right_sums: List[List[int]] = [[] for _ in range(n + 1)]

        # Generate all subsets for left half
        for count in range(n + 1):
            for combo in combinations(range(n), count):
                subset_sum = sum(left_half[i] for i in combo)
                left_sums[count].append(subset_sum)

        # Generate all subsets for right half
        for count in range(n + 1):
            for combo in combinations(range(n), count):
                subset_sum = sum(right_half[i] for i in combo)
                right_sums[count].append(subset_sum)

        # Sort right_sums for binary search
        for count in range(n + 1):
            right_sums[count].sort()

        min_diff = 1 << 31

        # Try all possible combinations
        # We take k elements from left and (n-k) elements from right
        for left_count in range(n + 1):
            right_count = n - left_count

            for left_sum in left_sums[left_count]:
                # We want: left_sum + right_sum ≈ target
                # So we need: right_sum ≈ target - left_sum
                needed = target - left_sum

                # Binary search for closest value in right_sums[right_count]
                right_list = right_sums[right_count]

                # Find closest value to 'needed'
                idx = bisect_left(right_list, needed)

                # Check idx and idx-1
                for i in [idx - 1, idx]:
                    if 0 <= i < len(right_list):
                        right_sum = right_list[i]
                        sum1 = left_sum + right_sum
                        sum2 = total - sum1
                        diff = abs(sum1 - sum2)
                        min_diff = min(min_diff, diff)

        return min_diff

    def test(self):
        for nums, expected in [
            ([3, 9, 7, 3], 2),
            ([-36, 36], 72),
            ([2, -1, 0, 4, -2, -9], 0),
        ]:
            output = self.minimumDifference(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
