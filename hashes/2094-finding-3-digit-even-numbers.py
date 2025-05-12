from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/finding-3-digit-even-numbers/

# python3 -m unittest hashes/2094-finding-3-digit-even-numbers.py

class Solution(unittest.TestCase):
    def findEvenNumbers(self, digits: List[int]) -> List[int]:
        count = Counter(digits)
        uniques = []
        for sadi in range(1,10):
            if count[sadi] == 0: continue
            count[sadi] -= 1
            for dahi in range(10):
                if count[dahi] == 0: continue
                count[dahi] -= 1
                for vohid in range(0, 10, 2):
                    if count[vohid] == 0: continue
                    count[vohid] -= 1
                    uniques.append(sadi*100 + dahi*10 + vohid)
                    count[vohid] += 1
                count[dahi] += 1
            count[sadi] += 1
        return uniques

    def test(self):
        for digits, expected in [
            ([2,1,3,0], [102,120,130,132,210,230,302,310,312,320]),
            ([2,2,8,8,2], [222,228,282,288,822,828,882]),
            ([3,7,5], []),
        ]:
            output = self.findEvenNumbers(digits)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
