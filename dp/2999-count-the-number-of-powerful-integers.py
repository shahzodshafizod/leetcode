import unittest

# TODO: Solve it by YOURSELF!!! # pylint: disable=fixme

# https://leetcode.com/problems/count-the-number-of-powerful-integers/

# python3 -m unittest dp/2999-count-the-number-of-powerful-integers.py


class Solution(unittest.TestCase):
    # Approach: Digital Dynamic Programming
    # Time: O(len(finish))
    # Space: O(len(finish))
    def numberOfPowerfulInt(self, start: int, finish: int, limit: int, s: str) -> int:
        def count(finish: int) -> int:
            flow = str(finish)
            n = len(flow)
            preflen = n - len(s)
            if preflen < 0:
                return 0
            dp = [[0] * 2 for _ in range(preflen + 1)]
            dp[preflen][0] = 1
            dp[preflen][1] = 1 if int(flow[preflen:]) >= int(s) else 0
            for idx in range(preflen - 1, -1, -1):
                dp[idx][0] = (limit + 1) * dp[idx + 1][0]
                if int(flow[idx]) <= limit:
                    dp[idx][1] = int(flow[idx]) * dp[idx + 1][0] + dp[idx + 1][1]
                else:
                    dp[idx][1] = (limit + 1) * dp[idx + 1][0]
            return dp[0][1]

        return count(finish) - count(start - 1)

    def test(self):
        for start, finish, limit, s, expected in [
            (1, 6000, 4, "124", 5),
            (15, 215, 6, "10", 2),
            (1000, 2000, 4, "3000", 0),
        ]:
            output = self.numberOfPowerfulInt(start, finish, limit, s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
