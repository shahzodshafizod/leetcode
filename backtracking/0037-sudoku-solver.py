from typing import List
import unittest

# python3 -m unittest backtracking/0037-sudoku-solver.py


class Solution(unittest.TestCase):
    def solveSudoku(self, board: List[List[str]]):
        # Check if it's valid and set existed values
        rows, cols, boxes = [0] * 9, [0] * 9, [0] * 9
        for row in range(9):
            for col in range(9):
                if board[row][col] == '.':
                    continue
                bit = 1 << int(board[row][col])
                bid = row // 3 * 3 + col // 3
                if rows[row] & bit or cols[col] & bit or boxes[bid] & bit:
                    return
                rows[row] ^= bit
                cols[col] ^= bit
                boxes[bid] ^= bit

        # Solve
        def backtracking(row: int, col: int) -> bool:
            if row == 9:
                return True
            col += 1
            if col == 9:
                return backtracking(row + 1, -1)
            if board[row][col] != '.':
                return backtracking(row, col)
            bid = row // 3 * 3 + col // 3
            for num in range(1, 10):
                bit = 1 << num
                if rows[row] & bit or cols[col] & bit or boxes[bid] & bit:
                    continue
                rows[row] ^= bit
                cols[col] ^= bit
                boxes[bid] ^= bit
                board[row][col] = str(num)
                if backtracking(row, col):
                    return True
                rows[row] ^= bit
                cols[col] ^= bit
                boxes[bid] ^= bit
                board[row][col] = '.'
            return False

        backtracking(0, -1)

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
                [
                    ["5", "3", "4", "6", "7", "8", "9", "1", "2"],
                    ["6", "7", "2", "1", "9", "5", "3", "4", "8"],
                    ["1", "9", "8", "3", "4", "2", "5", "6", "7"],
                    ["8", "5", "9", "7", "6", "1", "4", "2", "3"],
                    ["4", "2", "6", "8", "5", "3", "7", "9", "1"],
                    ["7", "1", "3", "9", "2", "4", "8", "5", "6"],
                    ["9", "6", "1", "5", "3", "7", "2", "8", "4"],
                    ["2", "8", "7", "4", "1", "9", "6", "3", "5"],
                    ["3", "4", "5", "2", "8", "6", "1", "7", "9"],
                ],
            ),
        ]:
            self.solveSudoku(board)
            output = board
            self.assertListEqual(output, expected, f"output: {output}, expected: {expected}")
