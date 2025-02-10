import unittest

# https://leetcode.com/problems/clear-digits/

# python3 -m unittest stacks/3174-clear-digits.py

class Solution(unittest.TestCase):
    def clearDigits(self, s: str) -> str:
        push, pop = (stack := []).append, stack.pop
        for c in s:
            if c.isdigit():
                pop()
            else:
                push(c)
        return "".join(stack)

    def test(self):
        for s, expected in [
            ("abc", "abc"),
            ("cb34", ""),
            ("a", "a"),
            ("xzuzr2ghilydk", "xzuzghilydk"),
            ("qm93xjkpaaovhqckjhg5j1o4rndtg3bobgeke", "xjkpaaovhqckjhrndtbobgeke"),
        ]:
            output = self.clearDigits(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
