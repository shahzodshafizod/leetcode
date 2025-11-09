from typing import List
import unittest

# https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/

# python3 -m unittest greedy/1326-minimum-number-of-taps-to-open-to-water-a-garden.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Bottom-Up Dynamic Programming)
    # # Time: O(nr), r=ranges[i]
    # # Space: O(n)
    # def minTaps(self, n: int, ranges: List[int]) -> int:
    #     dp = [n + 2] * (n + 1)
    #     dp[0] = 0
    #     for i, r in enumerate(ranges):
    #         left = max(0, i - r)
    #         right = min(n, i + r)
    #         for j in range(left, right + 1):
    #             dp[j] = min(dp[j], dp[left] + 1)
    #     return -1 if dp[n] == n + 2 else dp[n]

    # Approach #2: Greedy (Jump Game II)
    # Time: O(n)
    # Space: O(n)
    def minTaps(self, n: int, ranges: List[int]) -> int:
        maxReach = [0] * (n + 1)
        for i, r in enumerate(ranges):
            left = max(0, i - r)
            right = min(n, i + r)
            maxReach[left] = max(maxReach[left], right)

        taps = 0
        currentEnd = 0
        farthest = 0

        # Greedy: try to reach position n
        for i in range(n + 1):
            # If we can't reach this position, impossible
            if i > farthest:
                return -1

            # Update the farthest we can reach
            farthest = max(farthest, maxReach[i])

            # If we've reached the end of current tap's range
            if i == currentEnd:
                taps += 1
                currentEnd = farthest

                # If we can already reach the end, we're done
                if currentEnd >= n:
                    return taps

        return -1

    def test(self):
        for n, ranges, expected in [
            (5, [3, 4, 1, 1, 0, 0], 1),
            (3, [0, 0, 0, 0], -1),
        ]:
            output = self.minTaps(n, ranges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
