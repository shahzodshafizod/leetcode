import unittest

# https://leetcode.com/problems/n-queens-ii/

# python3 -m unittest old-submissions/hard/0052-n-queens-ii.py


class Solution(unittest.TestCase):
    def totalNQueens(self, n: int) -> int:
        cols = [False] * n
        diag = [False] * 2 * n  # [row+col] as slash /
        backDiag = [False] * 2 * n  # [n-row+col] as back slash \

        def backtrack(row: int) -> int:
            if row == n:
                return 1

            solutions = 0
            for col in range(n):
                if cols[col] or diag[row + col] or backDiag[n - row + col]:
                    continue

                cols[col] = True
                diag[row + col] = True
                backDiag[n - row + col] = True

                solutions += backtrack(row + 1)

                cols[col] = False
                diag[row + col] = False
                backDiag[n - row + col] = False

            return solutions

        return backtrack(0)

    def test(self):
        for n, expected in [
            (4, 2),
            (1, 1),
        ]:
            output = self.totalNQueens(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
