from typing import List
import unittest

# https://leetcode.com/problems/maximize-subarrays-after-removing-one-conflicting-pair/

# python3 -m unittest prefixsums/3480-maximize-subarrays-after-removing-one-conflicting-pair.py


class Solution(unittest.TestCase):
    # Time: O(n)
    # Space: O(n)
    def maxSubarrays(self, n: int, conflictingPairs: List[List[int]]) -> int:
        right: List[List[int]] = [[] for _ in range(n + 1)]
        for a, b in conflictingPairs:
            right[max(a, b)].append(min(a, b))
        res = 0
        left = [0, 0]
        after_removal = [0] * (n + 1)
        for r in range(1, n + 1):
            for l in right[r]:
                # left = max(left, [l, left[0]], [left[0], l])
                if l > left[0]:
                    left[1] = left[0]
                    left[0] = l
                elif l > left[1]:
                    left[1] = l
            res += r - left[0]
            after_removal[left[0]] += left[0] - left[1]
        return res + max(after_removal)

    def test(self):
        for n, conflictingPairs, expected in [
            (4, [[2, 3], [1, 4]], 9),
            (5, [[1, 2], [2, 5], [3, 5]], 12),
            (10, [[1, 2], [2, 3], [3, 5], [5, 7], [7, 9], [9, 10], [4, 6], [6, 8], [8, 10], [1, 10]], 18),
            (12, [[2, 4], [5, 8], [3, 9], [1, 12], [6, 11], [4, 10]], 51),
        ]:
            output = self.maxSubarrays(n, conflictingPairs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
