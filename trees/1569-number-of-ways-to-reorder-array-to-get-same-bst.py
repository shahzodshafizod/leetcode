from typing import List
import unittest

# https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/

# python3 -m unittest trees/1569-number-of-ways-to-reorder-array-to-get-same-bst.py

# # Type alias for BST structure
# BSTNode = Optional[Tuple[int, "BSTNode", "BSTNode"]]


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # Time: O(n! * n) - n! permutations, O(n) to build each BST - TLE for n > 10
    # # Space: O(n) - recursion depth
    # def numOfWays(self, nums: List[int]) -> int:
    #     MOD = 10**9 + 7

    #     def build_bst(arr: List[int]) -> BSTNode:
    #         if not arr:
    #             return None
    #         root = arr[0]
    #         left = [x for x in arr if x < root]
    #         right = [x for x in arr if x > root]
    #         return (root, build_bst(left), build_bst(right))

    #     original_bst = build_bst(nums)

    #     from itertools import permutations

    #     count = 0

    #     for perm in permutations(nums):
    #         if build_bst(list(perm)) == original_bst:
    #             count += 1

    #     # Subtract 1 for the original ordering
    #     return (count - 1) % MOD

    # Approach #2: Divide and Conquer with Combinatorics
    # Time: O(n^2) - computing binomial coefficients
    # Space: O(n^2) - storing Pascal's triangle
    def numOfWays(self, nums: List[int]) -> int:
        MOD = 10**9 + 7
        n = len(nums)

        # Precompute Pascal's triangle for binomial coefficients
        # pascal[i][j] = C(i, j)
        pascal = [[0] * (n + 1) for _ in range(n + 1)]
        for i in range(n + 1):
            pascal[i][0] = 1
            for j in range(1, i + 1):
                pascal[i][j] = (pascal[i - 1][j - 1] + pascal[i - 1][j]) % MOD

        def count_ways(arr: List[int]) -> int:
            if len(arr) <= 1:
                return 1

            root = arr[0]
            left = [x for x in arr if x < root]
            right = [x for x in arr if x > root]

            left_ways = count_ways(left)
            right_ways = count_ways(right)

            # Number of ways to interleave left and right
            # C(len(left) + len(right), len(left))
            ways = pascal[len(left) + len(right)][len(left)]

            return (ways * left_ways % MOD) * right_ways % MOD

        # Subtract 1 because the original ordering is not counted as a reordering
        return (count_ways(nums) - 1) % MOD

    def test(self):
        for nums, expected in [
            ([2, 1, 3], 1),  # [2, 3, 1] gives same BST
            ([3, 4, 5, 1, 2], 5),  # Multiple reorderings
            ([1, 2, 3], 0),  # Only one way (original)
            ([3, 1, 2, 5, 4, 6], 19),
            ([9, 4, 2, 1, 3, 6, 5, 7, 8, 14, 11, 10, 12, 13, 16, 15, 17, 18], 216212978),
        ]:
            output = self.numOfWays(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
