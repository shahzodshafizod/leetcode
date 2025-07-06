import unittest

# https://leetcode.com/problems/domino-and-tromino-tiling/

# python3 -m unittest dp/0790-domino-and-tromino-tiling.py

# Full[n] = Full[n-1] + Full[n-2] + TopMissing[n-1] + BottomMissing[n-1]

#   F[n-1]      F[n-2]        T[n-1]       B[n-1]
# |#######+|  |#######--|  |#######-+|  |########-|
# |#######+|  |#######--|  |########-|  |#######-+|


# TopMissing[n] = Full[n-2] + BottomMissing[n-1]
#   F[n-2]        B[n-1]
# |#######- |  |######## |
# |#######+-|  |#######--|


# BottomMissing[n] = Full[n-2] + TopMissing[n-1]
#   F[n-2]        T[n-1]
# |#######+-|  |#######--|
# |#######- |  |######## |


class Solution(unittest.TestCase):
    # # Approach: Bottom-Up Dynamic Programming
    # # Time: O(n)
    # # Space: O(n)
    # def numTilings(self, n: int) -> int:
    #     full = [0] * (n+1)
    #     tm = [0] * (n+1)
    #     bm = [0] * (n+1)
    #     full[0], full[1] = 1, 1
    #     for i in range(2, n+1):
    #         full[i] = full[i-1] + full[i-2] + tm[i-1] + bm[i-1]
    #         tm[i] = full[n-2] + bm[i-1]
    #         bm[i] = full[n-2] + tm[i-1]
    #     return full[n] % (10**9+7)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(n)
    # Space: O(1)
    def numTilings(self, n: int) -> int:
        tm, bm = 0, 0
        preprev, prev, curr = 1, 1, 1  # [n-2], [n-1], [n]
        for _ in range(2, n + 1):
            curr = curr + preprev + tm + bm
            tm, bm = preprev + bm, preprev + tm
            preprev, prev = prev, curr
        return curr % (10**9 + 7)

    def test(self):
        for n, expected in [
            (3, 5),
            (1, 1),
        ]:
            output = self.numTilings(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
