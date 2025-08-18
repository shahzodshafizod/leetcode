from typing import List
import unittest

# https://leetcode.com/problems/24-game/

# python3 -m unittest backtracking/0679-24-game.py


class Solution(unittest.TestCase):
    def judgePoint24(self, cards: List[int]) -> bool:
        def backtrack(cards: List[float]) -> bool:
            n = len(cards)
            if n == 1:
                return abs(cards[0] - 24) < 0.001
            for i in range(n):
                for j in range(i + 1, n):
                    c1 = cards[i]
                    c2 = cards[j]

                    values: List[float] = [c1 + c2, c1 - c2, c2 - c1, c1 * c2]
                    if c1:
                        values.append(c2 / c1)
                    if c2:
                        values.append(c1 / c2)

                    del cards[j]
                    del cards[i]

                    for value in values:
                        cards.append(value)
                        if backtrack(cards):
                            return True
                        cards.pop()
                    cards.insert(i, c1)
                    cards.insert(j, c2)
            return False

        # return backtrack([float(c) for c in cards])
        return backtrack(list(map(float, cards)))

    def test(self):
        for cards, expected in [
            ([4, 1, 8, 7], True),
            ([1, 2, 1, 2], False),
        ]:
            output = self.judgePoint24(cards)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
