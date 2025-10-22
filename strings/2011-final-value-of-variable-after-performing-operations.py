from typing import List
import unittest

# https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

# python3 -m unittest strings/2011-final-value-of-variable-after-performing-operations.py


class Solution(unittest.TestCase):
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        return sum(1 if op[1] == '+' else -1 for op in operations)

    def test(self):
        for operations, expected in [
            (["--X","X++","X++"], 1),
            (["++X","++X","X++"], 3),
            (["X++","++X","--X","X--"], 0),
        ]:
            output = self.finalValueAfterOperations(operations)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
