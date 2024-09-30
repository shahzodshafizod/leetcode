from typing import List
import unittest

# https://leetcode.com/problems/n-queens/

# python3 -m unittest backtracking/0051-n-queens.py

class Solution(unittest.TestCase):
    def solveNQueens(self, n: int) -> List[List[str]]:
        solutions = []
        board = [['.'] * n for _ in range(n)]
        cols = [False] * n # (col)
        negDiag = [False] * 2*n # (row-col)
        posDiag = [False] * 2*n # (row+col)
        def backtrack(row: int) -> None:
            if row == n:
                solutions.append(["".join(row) for row in board])
                return

            for col in range(n):
                if cols[col] or negDiag[n-row+col] or posDiag[row+col]:
                    continue

                board[row][col] = 'Q'
                cols[col] = True
                negDiag[n-row+col] = True
                posDiag[row+col] = True

                backtrack(row+1)

                board[row][col] = '.'
                cols[col] = False
                negDiag[n-row+col] = False
                posDiag[row+col] = False

        backtrack(0)
        return solutions

    def test(self) -> None:
        for n, expected in [
            (4, [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]),
            (1, [["Q"]]),
        ]:
            output = self.solveNQueens(n)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
