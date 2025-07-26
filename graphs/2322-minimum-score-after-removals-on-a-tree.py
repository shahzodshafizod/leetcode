from typing import List
import unittest

# https://leetcode.com/problems/minimum-score-after-removals-on-a-tree/

# python3 -m unittest graphs/2322-minimum-score-after-removals-on-a-tree.py


class Solution(unittest.TestCase):
    def minimumScore(self, nums: List[int], edges: List[List[int]]) -> int:
        n = len(nums)

        adj: List[List[int]] = [[] for _ in range(n)]
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)

        total = 0
        for num in nums:
            total ^= num

        tin = [0] * n  # time in
        tout = [0] * n  # time out
        node_xor = [0] * n

        def dfs(parent: int, node: int, timer: int) -> int:
            tin[node] = timer
            timer += 1
            node_xor[node] = nums[node]
            for nxt in adj[node]:
                if nxt != parent:
                    timer = dfs(node, nxt, timer)
                    node_xor[node] ^= node_xor[nxt]
            tout[node] = timer
            return timer + 1

        dfs(-1, 0, 0)

        def is_ancestor(node1: int, node2: int) -> bool:
            return tin[node1] < tin[node2] and tout[node1] > tout[node2]

        m = len(edges)
        res = (1 << 32) - 1
        for i in range(m - 1):
            a, b = edges[i]
            if is_ancestor(a, b):
                a, b = b, a

            for j in range(i + 1, m):
                c, d = edges[j]
                if is_ancestor(c, d):
                    c, d = d, c

                if is_ancestor(a, c):
                    x1 = total ^ node_xor[a]
                    x2 = node_xor[a] ^ node_xor[c]
                    x3 = node_xor[c]
                elif is_ancestor(c, a):
                    x1 = total ^ node_xor[c]
                    x2 = node_xor[c] ^ node_xor[a]
                    x3 = node_xor[a]
                else:
                    x1 = total ^ node_xor[a] ^ node_xor[c]
                    x2 = node_xor[a]
                    x3 = node_xor[c]

                res = min(res, max(x1, x2, x3) - min(x1, x2, x3))

        return res

    def test(self):
        for nums, edges, expected in [
            ([1, 5, 5, 4, 11], [[0, 1], [1, 2], [1, 3], [3, 4]], 9),
            ([5, 5, 2, 4, 4, 2], [[0, 1], [1, 2], [5, 2], [4, 3], [1, 3]], 0),
        ]:
            output = self.minimumScore(nums, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
