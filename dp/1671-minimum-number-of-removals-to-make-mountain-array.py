from typing import List
import unittest

# https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/

# python3 -m unittest dp/1671-minimum-number-of-removals-to-make-mountain-array.py


class Solution(unittest.TestCase):
    # # Approach: Dynamic Programming (Bottom-Up)
    # # Time: O(nn)
    # # Space: O(n)
    # def minimumMountainRemovals(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     lis = [1] * n
    #     for curr in range(n): # O(nn)
    #         for prev in range(curr):
    #             if nums[prev] < nums[curr]:
    #                 lis[curr] = max(lis[curr], 1+lis[prev])
    #     lds = [1] * n
    #     for curr in range(n-1,-1,-1): # O(nn)
    #         for next in range(curr+1, n):
    #             if nums[curr] > nums[next]:
    #                 lds[curr] = max(lds[curr], 1+lds[next])
    #     mountain = 0
    #     # minimum to remove = maximum to keep
    #     for peek in range(n): # O(n)
    #         length = lis[peek] + lds[peek] - 1
    #         if lis[peek] > 1 and lds[peek] > 1 and length >= 3:
    #             mountain = max(mountain, length)
    #     return n-mountain

    # Approach: Binary Search
    # Time: O(nlogn)
    # Space: O(n)
    def minimumMountainRemovals(self, nums: List[int]) -> int:
        n = len(nums)

        def get_lis(start: int, end: int, step: int) -> List[int]:
            lis = [0] * n
            sequence = []
            for idx in range(start, end, step):  # O(n)
                if not sequence or nums[idx] > sequence[-1]:
                    sequence.append(nums[idx])
                    lis[idx] = len(sequence)
                    continue
                left, right = 0, len(sequence)
                while left < right:  # O(logn)
                    mid = left + (right - left) // 2
                    if sequence[mid] < nums[idx]:
                        left = mid + 1
                    else:
                        right = mid
                sequence[left] = nums[idx]
                lis[idx] = left + 1
            return lis

        lis = get_lis(0, n, 1)
        lds = get_lis(n - 1, -1, -1)
        mountain = 0
        # minimum to remove = maximum to keep
        for peek in range(n):  # O(n)
            length = lis[peek] + lds[peek] - 1
            if lis[peek] > 1 and lds[peek] > 1 and length >= 3:
                mountain = max(mountain, length)
        return n - mountain

    def test(self):
        for nums, expected in [
            ([1, 3, 1], 0),
            ([2, 1, 1, 5, 6, 2, 3, 1], 3),
            ([9, 8, 1, 7, 6, 5, 4, 3, 2, 1], 2),
            ([100, 92, 89, 77, 74, 66, 64, 66, 64], 6),
            ([1, 16, 84, 9, 29, 71, 86, 79, 72, 12], 2),
            ([4, 5, 13, 17, 1, 7, 6, 11, 2, 8, 10, 15, 3, 9, 12, 14, 16], 10),
        ]:
            output = self.minimumMountainRemovals(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
