from collections import defaultdict, deque
import unittest
from typing import List

# https://leetcode.com/problems/jump-game-iv/

# python3 -m unittest graphs/1345-jump-game-iv.py

class Solution(unittest.TestCase):
    def minJumps(self, arr: List[int]) -> int:
        n = len(arr)
        lookup = defaultdict(list)
        for idx in range(n):
            lookup[arr[idx]].append(idx)
        queue = deque()
        visited = [False] * n
        queue.append((0, 0)) # index, jumps
        visited[0] = True
        while queue:
            curr, steps = queue.popleft()
            if curr == n - 1:
                return steps
            for next in [curr-1, curr+1]:
                if 0 <= next < n and not visited[next]:
                    queue.append((next, steps+1))
                    visited[next] = True
            for next in lookup[arr[curr]]:
                if not visited[next]:
                    queue.append((next, steps+1))
                    visited[next] = True
            lookup[arr[curr]].clear()
        return n-1

    def test(self):
        for arr, expected in [
            ([100,-23,-23,404,100,23,23,23,3,404], 3),
            ([7], 0),
            ([7,6,9,6,9,6,9,7], 1),
            ([6, 1, 9], 2),
            ([7, 7, 7, 7, 7, 1, 2, 3], 4),
            # ([7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,77,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,1,1,1,1,1,1,1,1,1,1,1,11], 4),
        ]:
            output = self.minJumps(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
