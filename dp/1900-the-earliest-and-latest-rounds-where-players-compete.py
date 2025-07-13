from typing import List
import unittest

# https://leetcode.com/problems/the-earliest-and-latest-rounds-where-players-compete/

# python3 -m unittest dp/1900-the-earliest-and-latest-rounds-where-players-compete.py


class Solution(unittest.TestCase):
    # # Approach #1: Bitmask + Dynamic Programming
    # # Time: O(2^n)
    # # Space: O(2^n)
    # def earliestAndLatest(self, n: int, firstPlayer: int, secondPlayer: int) -> List[int]:
    #     f, s = firstPlayer - 1, secondPlayer - 1
    #     earliest, latest = float("inf"), 0

    #     def dfs(left: int, right: int, mask: int, rnd: int):
    #         if left >= right:
    #             # We're done with the current round. Go to the next one
    #             dfs(0, n - 1, mask, rnd + 1)
    #         elif mask & (1 << left) == 0:
    #             # Skip the defeated left player
    #             dfs(left + 1, right, mask, rnd)
    #         elif mask & (1 << right) == 0:
    #             # Skip the defeated right player
    #             dfs(left, right - 1, mask, rnd)
    #         elif left == f and right == s:
    #             # Stop, f & s are fighting against each other
    #             nonlocal earliest, latest
    #             earliest = min(earliest, rnd)
    #             latest = max(latest, rnd)
    #         else:
    #             if left not in (f, s):
    #                 # Let left LOSE and right WIN
    #                 dfs(left + 1, right - 1, mask ^ (1 << left), rnd)
    #             if right not in (f, s):
    #                 # Let right LOSE and left WIN
    #                 dfs(left + 1, right - 1, mask ^ (1 << right), rnd)

    #     dfs(0, n - 1, (1 << n) - 1, 1)
    #     return [earliest, latest]

    # Approach #2: Counting + Bitmask + Dynamic Programming
    # Time: O(2^n)
    # Space: O(2^n)
    def earliestAndLatest(self, n: int, firstPlayer: int, secondPlayer: int) -> List[int]:
        f, s = firstPlayer - 1, secondPlayer - 1
        earliest, latest = float("inf"), 0
        seen = [[[False] * n for _ in range(n)] for _ in range(n)]

        def dfs(left: int, right: int, mask: int, rnd: int, lcnt: int, mcnt: int, rcnt: int):
            # lcnt: count of active players to the left
            # mcnt: count of active players in the middle
            # rcnt: count of active players to the right
            if left >= right:
                # We're done with the current round. Go to the next one
                dfs(0, n - 1, mask, rnd + 1, lcnt, mcnt, rcnt)
            elif mask & (1 << left) == 0:
                # Skip the defeated left player
                dfs(left + 1, right, mask, rnd, lcnt, mcnt, rcnt)
            elif mask & (1 << right) == 0:
                # Skip the defeated right player
                dfs(left, right - 1, mask, rnd, lcnt, mcnt, rcnt)
            elif left == f and right == s:
                # Stop, f & s are fighting against each other
                nonlocal earliest, latest
                earliest = min(earliest, rnd)
                latest = max(latest, rnd)
            elif not seen[lcnt][mcnt][rcnt]:
                seen[lcnt][mcnt][rcnt] = True
                if left not in (f, s):
                    # Let left LOSE and right WIN
                    dfs(
                        left + 1,
                        right - 1,
                        mask ^ (1 << left),
                        rnd,
                        lcnt - (1 if left < f else 0),
                        mcnt - (1 if f < left < s else 0),
                        rcnt - (1 if s < left else 0),
                    )
                if right not in (f, s):
                    # Let right LOSE and left WIN
                    dfs(
                        left + 1,
                        right - 1,
                        mask ^ (1 << right),
                        rnd,
                        lcnt - (1 if right < f else 0),
                        mcnt - (1 if f < right < s else 0),
                        rcnt - (1 if s < right else 0),
                    )

        dfs(0, n - 1, (1 << n) - 1, 1, f, s - f - 1, n - 1 - s)
        return [earliest, latest]

    def test(self):
        for n, firstPlayer, secondPlayer, expected in [
            (11, 2, 4, [3, 4]),
            (5, 1, 5, [1, 1]),
            (11, 4, 11, [2, 4]),
        ]:
            output = self.earliestAndLatest(n, firstPlayer, secondPlayer)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
