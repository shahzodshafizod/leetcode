from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/

# python3 -m unittest hashes/3160-find-the-number-of-distinct-colors-among-the-balls.py

class Solution(unittest.TestCase):
    def queryResults(self, limit: int, queries: List[List[int]]) -> List[int]:
        balls = {} # ball -> color
        colors = defaultdict(int) # color -> count
        result = [0] * len(queries)
        count = 0
        for idx, (x, y) in enumerate(queries):
            if x in balls:
                colors[balls[x]] -= 1
                if colors[balls[x]] == 0:
                    count -= 1
            balls[x] = y
            colors[balls[x]] += 1
            if colors[balls[x]] == 1:
                count += 1
            result[idx] = count
        return result

    def test(self):
        for limit, queries, expected in [
            (4, [[1,4],[2,5],[1,3],[3,4]], [1,2,2,3]),
            (4, [[0,1],[1,2],[2,2],[3,4],[4,5]], [1,2,2,3,4]),
        ]:
            output = self.queryResults(limit, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
