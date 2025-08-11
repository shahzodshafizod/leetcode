from typing import List
import unittest

# https://leetcode.com/problems/range-product-queries-of-powers/

# python3 -m unittest prefixsums/2438-range-product-queries-of-powers.py


class Solution(unittest.TestCase):
    def productQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        MOD = 10**9 + 7
        powers: List[int] = []
        for i in range(32):
            if (n >> i) & 1:
                powers.append(1 << i)
        m = len(powers)
        prepro: List[List[int]] = [[0] * m for _ in range(m)]
        for i in range(m):
            pro = 1
            for j in range(i, m):
                pro = pro * powers[j] % MOD
                prepro[i][j] = pro
        answer = [0] * len(queries)
        for i, (left, right) in enumerate(queries):
            answer[i] = prepro[left][right]
        return answer

    def test(self):
        for n, queries, expected in [
            (15, [[0, 1], [2, 2], [0, 3]], [2, 4, 64]),
            (2, [[0, 0]], [2]),
        ]:
            output = self.productQueries(n, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
