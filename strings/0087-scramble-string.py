import unittest
from typing import Dict, Tuple

# https://leetcode.com/problems/scramble-string/

# python3 -m unittest strings/0087-scramble-string.py


class Solution(unittest.TestCase):
    # Approach: Dynamic Programming with Memoization (Top-Down)
    # Time: O(n^4) - for each substring pair, try all split points
    # Space: O(n^3) - memoization cache
    def isScramble(self, s1: str, s2: str) -> bool:
        # Memoization cache: (s1, s2) -> bool
        memo: Dict[Tuple[str, str], bool] = {}
        def helper(s1: str, s2: str) -> bool:
            # Base cases
            if s1 == s2:
                return True

            # Quick check: if character frequencies don't match, can't be scramble
            if sorted(s1) != sorted(s2):
                return False

            # Check memo
            key = (s1, s2)
            if key in memo:
                return memo[key]

            # Try all possible split points
            n = len(s1)
            for i in range(1, n):
                # Case 1: No swap
                # s1[:i] matches s2[:i] AND s1[i:] matches s2[i:]
                if helper(s1[:i], s2[:i]) and helper(s1[i:], s2[i:]):
                    memo[key] = True
                    return True

                # Case 2: Swap
                # s1[:i] matches s2[n-i:] AND s1[i:] matches s2[:n-i]
                if helper(s1[:i], s2[n-i:]) and helper(s1[i:], s2[:n-i]):
                    memo[key] = True
                    return True

            memo[key] = False
            return False

        return helper(s1, s2)

    def test(self):
        for s1, s2, expected in [
            ("great", "rgeat", True),
            ("abcde", "caebd", False),
            ("a", "a", True),
            ("abc", "bca", True),
            ("abcdbdacbdac", "bdacabcdbdac", True),
        ]:
            output = self.isScramble(s1, s2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
