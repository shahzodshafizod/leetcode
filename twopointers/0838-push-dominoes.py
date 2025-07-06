import unittest

# https://leetcode.com/problems/push-dominoes/

# python3 -m unittest twopointers/0838-push-dominoes.py


class Solution(unittest.TestCase):
    # # Approach: Calculate Force
    # # Time: O(n)
    # # Space: O(n)
    # def pushDominoes(self, dominoes: str) -> str:
    #     n = len(dominoes)
    #     force = [0]*n
    #     f = 0
    #     for idx in range(n):
    #         c = dominoes[idx]
    #         f = n if c == 'R' else 0 if c == 'L' else max(f-1, 0)
    #         force[idx] = f
    #     f = 0
    #     for idx in range(n-1,-1,-1):
    #         c = dominoes[idx]
    #         f = n if c == 'L' else 0 if c == 'R' else max(f-1, 0)
    #         force[idx] -= f
    #     return "".join('.' if f == 0 else 'R' if f > 0 else 'L' for f in force)

    # Approach: Replacing
    # Time: O(n)
    # Space: O(n)
    def pushDominoes(self, dominoes: str) -> str:
        tmp = ''
        while tmp != dominoes:
            tmp = dominoes
            dominoes = dominoes.replace("R.L", "###")
            dominoes = dominoes.replace("R.", "RR")
            dominoes = dominoes.replace(".L", "LL")
        return dominoes.replace("###", "R.L")

    def test(self):
        for dominoes, expected in [
            ("RR.L", "RR.L"),
            (".L.R...LR..L..", "LL.RR.LLRRLL.."),
        ]:
            output = self.pushDominoes(dominoes)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
