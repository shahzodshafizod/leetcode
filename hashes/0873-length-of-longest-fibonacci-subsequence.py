from typing import List
import unittest

# https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/

# python3 -m unittest hashes/0873-length-of-longest-fibonacci-subsequence.py

class Solution(unittest.TestCase):
    # Approach: Hash
    # Time: O(nn logm), n=len(arr), m=len(fib seq)
    # Space: O(n)
    def lenLongestFibSubseq(self, arr: List[int]) -> int:
        nums = set(arr)
        n = len(arr)
        longest = 0
        for i in range(n):
            for j in range(i+1, n):
                prev, curr = arr[i], arr[j]
                length = 2
                while prev+curr in nums:
                    length += 1
                    longest = max(longest, length)
                    prev, curr = curr, prev+curr
        return longest

    def test(self):
        for arr, expected in [
            ([1,2,3,4,5,6,7,8], 5),
            ([1,3,7,11,12,14,18], 3),
        ]:
            output = self.lenLongestFibSubseq(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
