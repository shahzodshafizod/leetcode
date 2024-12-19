import unittest

class Solution(unittest.TestCase):
    def a(self):
        pass

    def test(self):
        for expected in [
            # (),
        ]:
            output = self.a()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
