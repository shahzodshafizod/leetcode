import unittest

# https://leetcode.com/problems/water-bottles/

# python3 -m unittest maths/1518-water-bottles.py


class Solution(unittest.TestCase):
    # Approach: Simulation
    # Time: O(logm(n))
    # Space: O(1)
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        # return numBottles + (numBottles - 1) // (numExchange - 1)
        consumed, empty = 0, 0
        while numBottles:
            consumed += numBottles
            empty += numBottles
            numBottles = empty // numExchange
            empty %= numExchange
        return consumed

    def test(self):
        for numBottles, numExchange, expected in [
            (9, 3, 13),
            (15, 4, 19),
        ]:
            output = self.numWaterBottles(numBottles, numExchange)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
