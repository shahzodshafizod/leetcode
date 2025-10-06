from typing import List
import unittest

# https://leetcode.com/problems/similar-string-groups/

# python3 -m unittest graphs/0839-similar-string-groups.py


class Solution(unittest.TestCase):
    # # Approach #1: Union Find
    # # Time: O(nnm), n=len(strs), m=len(strs[i])
    # # Space: O(nm)
    # def numSimilarGroups(self, strs: List[str]) -> int:
    #     n = len(strs)
    #     root = list(range(n))
    #     groups = n

    #     def find(x: int) -> int:
    #         if root[x] != x:
    #             root[x] = find(root[x])
    #         return root[x]

    #     def union(x: int, y: int):
    #         nonlocal groups
    #         rx, ry = find(x), find(y)
    #         if rx < ry:
    #             root[ry] = rx
    #             groups -= 1
    #         elif rx > ry:
    #             root[rx] = ry
    #             groups -= 1

    #     def similar(i: int, j: int) -> bool:
    #         cnt = 0
    #         for k in range(len(strs[i])):
    #             if strs[i][k] != strs[j][k]:
    #                 cnt += 1
    #         return cnt in (2, 0)

    #     for i in range(n):
    #         for j in range(i + 1, n):
    #             if similar(i, j):
    #                 union(i, j)

    #     return groups

    # Approach #2: Depth-First Search
    # Time: O(nnm), n=len(strs), m=len(strs[i])
    # Space: O(nm)
    def numSimilarGroups(self, strs: List[str]) -> int:
        n = len(strs)
        seen = [False] * n
        groups = 0

        def similar(i: int, j: int) -> bool:
            cnt, m, k = 0, len(strs[i]), 0
            while k < m and cnt <= 2:
                if strs[i][k] != strs[j][k]:
                    cnt += 1
                k += 1
            return cnt <= 2

        def dfs(i: int):
            seen[i] = True
            for j in range(n):
                if not seen[j] and similar(i, j):
                    dfs(j)

        for i in range(n):
            if not seen[i]:
                groups += 1
                dfs(i)

        return groups

    def test(self):
        for strs, expected in [
            (["tars", "rats", "arts", "star"], 2),
            (["omv", "ovm"], 1),
            (["kccomwcgcs", "socgcmcwkc", "sgckwcmcoc", "coswcmcgkc", "cowkccmsgc", "cosgmccwkc", "sgmkwcccoc", "coswmccgkc", "kowcccmsgc", "kgcomwcccs"], 5),
        ]:
            output = self.numSimilarGroups(strs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
