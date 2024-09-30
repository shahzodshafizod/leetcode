from functools import cache
from typing import List
import unittest

"""
There are nCk(m + n, m) = nCk(m + n, n) solutions
to go from (0, 0) to (m, n)
because we have to perform m moves down and n moves right

First we move one step left (col-1) and compute the possibilities
nCk(m + (n-1), m)
if k <= nCk(m + (n-1), m), then the move will be left-right: H (horizontal)
else, we move up-down: V (vertical) and we decrease the number of ways: k -= c
"""

# https://leetcode.com/problems/kth-smallest-instructions/

# python3 -m unittest maths/1643-kth-smallest-instructions.py

class Solution(unittest.TestCase):
    # # Approach: Brute-Force (Memory+Time Limit Exceeded Exceptions)
    # def kthSmallestPath(self, destination: List[int], k: int) -> str:
    #     rdest, cdest = destination
    #     @cache
    #     def dfs(row: int, col: int, path: str) -> str:
    #         if row > rdest or col > cdest:
    #             return ""
    #         if row == rdest and col == cdest:
    #             nonlocal k
    #             k -= 1
    #             if k == 0:
    #                 return path
    #             return ""
    #         for r,c,p in [(row,col+1,path+'H'),(row+1,col,path+'V')]:
    #             result = dfs(r, c, p)
    #             if result != "":
    #                 return result
    #         return ""
    #     return dfs(0, 0, "")

    # Credit to: http://ptaskbook.com/en/tasks/recur.php (Recur6)
    dp = [[None] * 30 for _ in range(30)]
    def comb(self, n: int, k: int) -> int:
        if n == k or k == 0:
            return 1
        if self.dp[n][k] != None:
            return self.dp[n][k]
        self.dp[n][k] = self.comb(n-1, k) + self.comb(n-1, k-1)
        return self.dp[n][k]

    # # Approach: Recursive
    # # Time: O((row+col)^2)
    # # Space: O(row+col)
    # def kthSmallestPath(self, destination: List[int], k: int) -> str:
    #     row, col = destination
    #     if row == 0: return 'H'*col
    #     if col == 0: return 'V'*row
    #     c = self.comb(row+col-1, row)
    #     if k > c:
    #         return 'V' + self.kthSmallestPath([row-1, col], k-c)
    #     return 'H' + self.kthSmallestPath([row, col-1], k)

    # Approach: Iterative
    # Time: O((row+col)^2)
    # Space: O(1)
    def kthSmallestPath(self, destination: List[int], k: int) -> str:
        row, col = destination
        path = ""
        while row and col:
            c = self.comb(row+col-1, row)
            if k <= c:
                path += 'H'
                col -= 1
            else:
                k -= c
                path += 'V'
                row -= 1
        path += 'H'*col + 'V'*row
        return path

    def test(self) -> None:
        for destination, k, expected in [
            ([2,3], 1, "HHHVV"),
            ([2,3], 2, "HHVHV"),
            ([2,3], 3, "HHVVH"),
            ([15,15], 155117520, "VVVVVVVVVVVVVVVHHHHHHHHHHHHHHH"),
        ]:
            output = self.kthSmallestPath(destination, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
