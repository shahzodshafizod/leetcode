import unittest

# https://leetcode.com/problems/alice-and-bob-playing-flower-game/

# python3 -m unittest maths/3021-alice-and-bob-playing-flower-game.py


class Solution(unittest.TestCase):
    def flowerGame(self, n: int, m: int) -> int:
        # nev, nod = n // 2, (n + 1) // 2
        # mev, mod = m // 2, (m + 1) // 2
        # return (nev * mod) + (nod * mev)
        return m * n // 2

    def test(self):
        for n, m, expected in [
            (3, 2, 3),
            (1, 1, 0),
        ]:
            output = self.flowerGame(n, m)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
