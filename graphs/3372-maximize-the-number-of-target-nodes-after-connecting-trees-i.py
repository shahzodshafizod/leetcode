from typing import List
import unittest

# https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/

# python3 -m unittest graphs/3372-maximize-the-number-of-target-nodes-after-connecting-trees-i.py


class Solution(unittest.TestCase):
    # Approach: Depth-First Search
    # Time: O(nn+mm)
    # Space: O(n+m)
    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]], k: int) -> List[int]:
        def dfs(parent: int, node: int, adj: List[List[int]], k: int) -> int:
            if k < 0:
                return 0
            cnt = 1
            for nei in adj[node]:
                if nei != parent:
                    cnt += dfs(node, nei, adj, k - 1)
            return cnt

        def traverse(edges: List[List[int]], k: int, val: int) -> List[int]:
            n = len(edges) + 1
            adj = [[] for _ in range(n)]
            for a, b in edges:
                adj[a].append(b)
                adj[b].append(a)
            answer = [val] * n
            for node in range(n):
                answer[node] += dfs(-1, node, adj, k)
            return answer

        return traverse(edges1, k, max(traverse(edges2, k - 1, 0)))

    def test(self):
        for edges1, edges2, k, expected in [
            (
                [[0, 1], [0, 2], [2, 3], [2, 4]],
                [[0, 1], [0, 2], [0, 3], [2, 7], [1, 4], [4, 5], [4, 6]],
                2,
                [9, 7, 9, 8, 8],
            ),
            ([[0, 1], [0, 2], [0, 3], [0, 4]], [[0, 1], [1, 2], [2, 3]], 1, [6, 3, 3, 3, 3]),
        ]:
            output = self.maxTargetNodes(edges1, edges2, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
