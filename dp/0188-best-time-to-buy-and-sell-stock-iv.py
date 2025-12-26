from typing import List
import unittest

# https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/

# python3 -m unittest dp/0188-best-time-to-buy-and-sell-stock-iv.py


class Solution(unittest.TestCase):
    # Approach: Dynamic Programming
    # Time: O(n*k), Space: O(k)
    def maxProfit(self, k: int, prices: List[int]) -> int:
        n = len(prices)
        if not prices or k == 0:
            return 0

        # Optimization: unlimited transactions
        if k >= n // 2:
            return sum(max(prices[i] - prices[i - 1], 0) for i in range(1, n))

        buy = [-prices[0]] * (k + 1)
        sell = [0] * (k + 1)

        for i in range(1, n):
            for j in range(k, 0, -1):
                sell[j] = max(sell[j], buy[j] + prices[i])
                buy[j] = max(buy[j], sell[j - 1] - prices[i])

        return sell[k]

    def test(self):
        for k, prices, expected in [
            (2, [2, 4, 1], 2),
            (2, [3, 2, 6, 5, 0, 3], 7),
            (2, [3, 3, 5, 0, 0, 3, 1, 4], 6),
        ]:
            output = self.maxProfit(k, prices)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
