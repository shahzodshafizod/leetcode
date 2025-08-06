from typing import List
import unittest

# https://leetcode.com/problems/fruits-into-baskets-iii/

# python3 -m unittest intervals/3479-fruits-into-baskets-iii.py


class Solution(unittest.TestCase):
    def numOfUnplacedFruits(self, fruits: List[int], baskets: List[int]) -> int:
        n = len(baskets)
        segtree = [0] * (4 * n)

        def build(node: int, left: int, right: int):
            if left == right:
                segtree[node] = baskets[left]
                return
            mid = (left + right) // 2
            build(2 * node, left, mid)
            build(2 * node + 1, mid + 1, right)
            segtree[node] = max(segtree[2 * node], segtree[2 * node + 1])

        build(1, 0, n - 1)

        def find(node: int, fruit: int, left: int, right: int) -> bool:
            if fruit > segtree[node]:
                return False
            if left == right:
                segtree[node] = 0
                return True
            mid = (left + right) // 2
            res = find(2 * node, fruit, left, mid)
            if not res:
                res = find(2 * node + 1, fruit, mid + 1, right)
            segtree[node] = max(segtree[2 * node], segtree[2 * node + 1])
            return res

        res = 0
        for fruit in fruits:
            res += 1 if not find(1, fruit, 0, n - 1) else 0
        return res

    def test(self):
        for fruits, baskets, expected in [
            ([4, 2, 5], [3, 5, 4], 1),
            ([3, 6, 1], [6, 4, 7], 0),
        ]:
            output = self.numOfUnplacedFruits(fruits, baskets)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
