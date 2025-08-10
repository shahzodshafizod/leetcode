from typing import List
import unittest

# https://leetcode.com/problems/count-valid-paths-in-a-tree/

# python3 -m unittest unionfinds/2867-count-valid-paths-in-a-tree.py


class Solution(unittest.TestCase):
    def countPaths(self, n: int, edges: List[List[int]]) -> int:
        primes: List[bool] = [True] * (n + 1)
        primes[0] = primes[1] = False
        for num in range(2, int(n**0.5) + 1):
            if primes[num]:
                for nxt in range(2 * num, n + 1, num):
                    primes[nxt] = False

        graph: List[List[int]] = [[] for _ in range(n + 1)]
        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        roots = list(range(n + 1))

        def find(x: int) -> int:
            if roots[x] != x:
                roots[x] = find(roots[x])
            return roots[x]

        def union(u: int, v: int):
            ru = find(u)
            rv = find(v)
            roots[min(ru, rv)] = max(ru, rv)

        def dfs(parent: int, node: int) -> int:
            count = 1
            for nxt in graph[node]:
                if nxt != parent and not primes[nxt]:
                    union(node, nxt)
                    count += dfs(node, nxt)
            return count

        res = 0
        size = [0] * (n + 1)
        for node in range(1, n + 1):
            if primes[node]:
                total = 0
                for nxt in graph[node]:
                    if not primes[nxt]:
                        if size[find(nxt)] == 0:
                            size[find(nxt)] = dfs(node, nxt)
                        total += size[find(nxt)]
                res += total
                for nxt in graph[node]:
                    if not primes[nxt]:
                        total -= size[find(nxt)]
                        res += size[find(nxt)] * total

        return res

    def test(self):
        for n, edges, expected in [
            (5, [[1, 2], [1, 3], [2, 4], [2, 5]], 4),
            (6, [[1, 2], [1, 3], [2, 4], [3, 5], [3, 6]], 6),
            (4, [[1, 2], [4, 1], [3, 4]], 4),
        ]:
            output = self.countPaths(n, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
