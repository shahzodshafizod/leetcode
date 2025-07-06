import unittest

# https://leetcode.com/problems//

# python3 -m unittest template/template.py


class Solution(unittest.TestCase):
    def a(self) -> int:
        return 0

    def test(self):
        for expected in [
            # (),
        ]:
            output = self.a()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
