import unittest

# https://leetcode.com/problems/robot-return-to-origin/

# python3 -m unittest strings/0657-robot-return-to-origin.py


class Solution(unittest.TestCase):
    def judgeCircle(self, moves: str) -> bool:
        return moves.count('R') == moves.count('L') and moves.count('U') == moves.count('D')

    def test(self):
        for moves, expected in [
            ("UD", True),
            ("LL", False),
        ]:
            output = self.judgeCircle(moves)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
