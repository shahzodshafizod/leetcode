import unittest

# https://leetcode.com/problems/nim-game/

# python3 -m unittest maths/0292-nim-game.py

# n=1: Win (take 1 stone)
# n=2: Win (take 2 stones)
# n=3: Win (take 3 stones)
# n=4: Lose! No matter what we take:
#      - Take 1 → opponent faces 3 (wins)
#      - Take 2 → opponent faces 2 (wins)
#      - Take 3 → opponent faces 1 (wins)
# n=5: Win (take 1, leave 4 for opponent)
# n=6: Win (take 2, leave 4 for opponent)
# n=7: Win (take 3, leave 4 for opponent)
# n=8: Lose (same pattern as 4)


class Solution(unittest.TestCase):
    # # Approach #1: Use DP to determine winning positions
    # # Time: O(n)
    # # Space: O(n) for memoization
    # def canWinNim(self, n: int) -> bool:
    #     # dp[i] represents whether we can win with i stones
    #     dp = [False] * (n + 1)
    #     dp[1] = dp[2] = dp[3] = True

    #     for i in range(4, n + 1):
    #         # We win if we can make a move that leaves opponent in losing position
    #         # We can take 1, 2, or 3 stones
    #         dp[i] = not dp[i - 1] or not dp[i - 2] or not dp[i - 3]

    #     return dp[n]

    # Approach #2: Mathematical Pattern Recognition
    # Time: O(1)
    # Space: O(1)
    def canWinNim(self, n: int) -> bool:
        return n % 4 != 0

    def test(self):
        for n, expected in [
            (1, True),
            (2, True),
            (3, True),
            (4, False),
            (5, True),
            (6, True),
            (7, True),
            (8, False),
            (12, False),
            (15, True),
        ]:
            output = self.canWinNim(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
