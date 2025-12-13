import unittest

# https://leetcode.com/problems/smallest-substring-with-identical-characters-i/

# python3 -m unittest binarysearch/3398-smallest-substring-with-identical-characters-i.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Try all possible lengths from 1 to n
    # # Time: O(n^2) - for each length, validate in O(n)
    # # Space: O(1)
    # def minLength(self, s: str, numOps: int) -> int:
    #     n = len(s)

    #     # Special case: alternating pattern
    #     pattern1, pattern2 = 0, 0
    #     for i, c in enumerate(list(map(int, s))):
    #         pattern1 += 1 if c != i % 2 else 0
    #         pattern2 += 1 if c != 1 - i % 2 else 0
    #     if min(pattern1, pattern2) <= numOps:
    #         return 1

    #     def canAchieve(maxLen: int) -> bool:
    #         ops, i = 0, 0
    #         while i < n:
    #             j = i
    #             while j < n and s[j] == s[i]:
    #                 j += 1
    #             segmentLen = j - i

    #             if segmentLen > maxLen:
    #                 ops += segmentLen // (maxLen + 1)

    #             i = j
    #         return ops <= numOps

    #     for length in range(2, n + 1):
    #         if canAchieve(length):
    #             return length

    #     return n

    # Approach #2: Binary Search on Answer
    # Time: O(n * logn)
    # Space: O(1)
    def minLength(self, s: str, numOps: int) -> int:
        n = len(s)

        # Special case: alternating pattern
        # Try both "010101..." and "101010..."
        pattern1, pattern2 = 0, 0
        for i, c in enumerate(list(map(int, s))):
            pattern1 += 1 if c != i % 2 else 0
            pattern2 += 1 if c != 1 - i % 2 else 0
        if min(pattern1, pattern2) <= numOps:
            return 1

        def canAchieve(maxLen: int) -> bool:
            ops, i = 0, 0
            while i < n:
                j = i
                while j < n and s[j] == s[i]:
                    j += 1
                segmentLen = j - i
                if segmentLen > maxLen:
                    ops += segmentLen // (maxLen + 1)
                i = j
            return ops <= numOps

        left, right = 2, n
        while left < right:
            mid = (left + right) // 2
            if canAchieve(mid):
                right = mid
            else:
                left = mid + 1

        return right

    def test(self):
        for s, numOps, expected in [
            ("000001", 1, 2),
            ("0000", 2, 1),
            ("0101", 0, 1),
            ("1111", 1, 2),
            ("00", 0, 2),
            ("01", 0, 1),
            ("000", 1, 1),
        ]:
            output = self.minLength(s, numOps)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
