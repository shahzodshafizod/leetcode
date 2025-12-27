import unittest
from typing import List

# https://leetcode.com/problems/maximum-product-of-three-numbers/

# python3 -m unittest maths/0628-maximum-product-of-three-numbers.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Try all combinations)
    # # Time: O(n^3) - three nested loops
    # # Space: O(1) - constant space
    # def maximumProduct(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     max_product = -(1 << 31)
    
    #     # Try all possible triplets
    #     for i in range(n):
    #         for j in range(i + 1, n):
    #             for k in range(j + 1, n):
    #                 product = nums[i] * nums[j] * nums[k]
    #                 max_product = max(max_product, product)
    
    #     return max_product

    # Approach #2: Sorting + Math
    # Time: O(n log n) - sorting dominates
    # Space: O(1) - constant space (ignoring sort space)
    def maximumProduct(self, nums: List[int]) -> int:
        # Sort the array
        nums.sort()
        n = len(nums)

        # Maximum product can be either:
        # 1. Product of three largest numbers (all positive or all negative)
        # 2. Product of two smallest (most negative) and largest (if negatives exist)

        # Case 1: Three largest numbers
        product1 = nums[n - 1] * nums[n - 2] * nums[n - 3]

        # Case 2: Two smallest (negative) * largest (positive)
        product2 = nums[0] * nums[1] * nums[n - 1]

        return max(product1, product2)

    def test(self):
        for nums, expected in [
            ([1, 2, 3], 6),
            ([1, 2, 3, 4], 24),
            ([-1, -2, -3], -6),
            ([-4, -3, -2, -1, 60], 720),  # -4 * -3 * 60 = 720
            ([1, 2, 3, 4, 5, 6], 120),
            ([-10, -10, 1, 3, 2], 300),  # -10 * -10 * 3 = 300
        ]:
            output = self.maximumProduct(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
