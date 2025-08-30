from typing import List
import unittest

# https://leetcode.com/problems/valid-sudoku/

# python3 -m unittest bits/0036-valid-sudoku.py


class Solution(unittest.TestCase):
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows, cols, boxes = [0] * 9, [0] * 9, [0] * 9
        for row in range(9):
            for col in range(9):
                if board[row][col] == '.':
                    continue
                idx = row // 3 * 3 + col // 3
                bit = 1 << int(board[row][col])
                if (rows[row] & bit) or (cols[col] & bit) or (boxes[idx] & bit):
                    return False
                rows[row] |= bit
                cols[col] |= bit
                boxes[idx] |= bit
        return True

    def test(self):
        for board, expected in [
            (
                [
                    ["5", "3", ".", ".", "7", ".", ".", ".", "."],
                    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
                    [".", "9", "8", ".", ".", ".", ".", "6", "."],
                    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
                    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
                    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
                    [".", "6", ".", ".", ".", ".", "2", "8", "."],
                    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
                    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
                ],
                True,
            ),
            (
                [
                    ["8", "3", ".", ".", "7", ".", ".", ".", "."],
                    ["6", ".", ".", "1", "9", "5", ".", ".", "."],
                    [".", "9", "8", ".", ".", ".", ".", "6", "."],
                    ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
                    ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
                    ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
                    [".", "6", ".", ".", ".", ".", "2", "8", "."],
                    [".", ".", ".", "4", "1", "9", ".", ".", "5"],
                    [".", ".", ".", ".", "8", ".", ".", "7", "9"],
                ],
                False,
            ),
        ]:
            output = self.isValidSudoku(board)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
