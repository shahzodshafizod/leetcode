from typing import List
import unittest

# https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/

# python3 -m unittest bits/2657-find-the-prefix-common-array-of-two-arrays.py

class Solution(unittest.TestCase):
    # # Approach: Counting
    # # Time: O(n)
    # # Space: O(n)
    # def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
    #     n = len(A)
    #     C = [0] * n
    #     count = [0] * (n+1)
    #     answer = 0
    #     for idx in range(n):
    #         count[A[idx]] += 1
    #         if count[A[idx]] == 2:
    #             answer += 1
    #         count[B[idx]] += 1
    #         if count[B[idx]] == 2:
    #             answer += 1
    #         C[idx] = answer
    #     return C

    # Approach: Bit Manipulation
    # Time: O(n)
    # Space: O(1)
    def findThePrefixCommonArray(self, A: List[int], B: List[int]) -> List[int]:
        n = len(A)
        C = [0] * n
        amask, bmask = 0, 0
        for idx in range(n):
            amask |= 1 << A[idx]
            bmask |= 1 << B[idx]
            C[idx] = (amask & bmask).bit_count()
        return C
    
    def test(self):
        for A, B, expected in [
            ([1,3,2,4], [3,1,2,4], [0,2,3,4]),
            ([2,3,1], [3,1,2], [0,1,3]),
        ]:
            output = self.findThePrefixCommonArray(A, B)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
