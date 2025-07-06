from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/snakes-and-ladders/

# python3 -m unittest matrices/0909-snakes-and-ladders.py


class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # Time: O(nn)
    # Space: O(nn)
    def snakesAndLadders(self, board: List[List[int]]) -> int:
        n = len(board)

        def get_pos(square: int) -> tuple[int, int]:
            row = (square - 1) // n
            col = (square - 1) % n
            if row & 1:
                col = n - 1 - col
            return (n - 1 - row, col)

        target = n * n
        visited = [False] * target
        queue = deque([(1, 0)])  # (square, moves)
        while queue:
            square, moves = queue.popleft()
            moves += 1
            for nei in range(min(square + 6, target), square, -1):
                row, col = get_pos(nei)
                if board[row][col] != -1:
                    nei = board[row][col]
                if nei == target:
                    return moves
                if not visited[nei - 1]:
                    visited[nei - 1] = True
                    queue.append((nei, moves))
        return -1

    def test(self):
        for board, expected in [
            (
                [
                    [-1, -1, -1, -1, -1, -1],
                    [-1, -1, -1, -1, -1, -1],
                    [-1, -1, -1, -1, -1, -1],
                    [-1, 35, -1, -1, 13, -1],
                    [-1, -1, -1, -1, -1, -1],
                    [-1, 15, -1, -1, -1, -1],
                ],
                4,
            ),
            ([[-1, -1], [-1, 3]], 1),
        ]:
            output = self.snakesAndLadders(board)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
