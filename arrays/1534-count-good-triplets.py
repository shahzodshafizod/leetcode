from typing import List
import unittest

# https://leetcode.com/problems/count-good-triplets/

# python3 -m unittest arrays/1534-count-good-triplets.py

class Solution(unittest.TestCase):
    # Approach: Enumeration
    # Time: O(nnn)
    # Space: O(1)
    def countGoodTriplets(self, arr: List[int], a: int, b: int, c: int) -> int:
        count, n = 0, len(arr)
        for i in range(n-2):
            x = arr[i]
            for j in range(i+1, n-1):
                y = arr[j]
                if abs(x-y) > a: continue
                for k in range(j+1, n):
                    z = arr[k]
                    count += 1 if abs(y-z)<=b and abs(z-x)<=c else 0
        return count

    def test(self):
        for arr, a, b, c, expected in [
            ([3,0,1,1,9,7], 7, 2, 3, 4),
            ([1,1,2,2,3], 0, 0, 1, 0),
        ]:
            output = self.countGoodTriplets(arr, a, b, c)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
