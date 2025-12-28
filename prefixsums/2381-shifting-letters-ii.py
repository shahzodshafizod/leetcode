from typing import List
import unittest

# https://leetcode.com/problems/shifting-letters-ii/

# python3 -m unittest prefixsums/2381-shifting-letters-ii.py


class Solution(unittest.TestCase):
    # # Approach: Heap (Priority Queue) (Time Limit Exceeded)
    # # Time: O(n log m), n=len(s), m=len(shifts)
    # # Space: O(m)
    # def shiftingLetters(self, s: str, shifts: List[List[int]]) -> str:
    #     heapq.heapify(shifts)
    #     s = list(s)
    #     for idx in range(len(s)):
    #         if not shifts: break
    #         code = ord(s[idx]) - ord('a')
    #         while shifts and shifts[0][0] == idx:
    #             start, end, direction = shifts[0]
    #             if direction == 1:
    #                 code += 1
    #             else:
    #                 code -= 1
    #             code = (code + 26) % 26
    #             start += 1
    #             if start <= end:
    #                 heapq.heapreplace(shifts, [start, end, direction])
    #             else:
    #                 heapq.heappop(shifts)
    #         s[idx] = chr(ord('a') + code)
    #     return "".join(s)

    # Approach: Line Sweep
    # Time: O(n+m), n=len(s), m=len(shifts)
    # Space: O(n)
    def shiftingLetters(self, s: str, shifts: List[List[int]]) -> str:
        line = [0] * (len(s) + 1)
        for start, end, direction in shifts:
            if direction == 0:
                direction = -1
            line[start] += direction
            line[end + 1] -= direction
        sc = list(s)
        delta = 0
        for idx, c in enumerate(sc):
            delta += line[idx]
            code = (ord(c) - ord('a') + delta) % 26
            if code < 0:
                code += 26
            sc[idx] = chr(ord('a') + code)
        return "".join(sc)

    def test(self):
        for s, shifts, expected in [
            ("abc", [[0, 1, 0], [1, 2, 1], [0, 2, 1]], "ace"),
            ("dztz", [[0, 0, 0], [1, 1, 1]], "catz"),
        ]:
            output = self.shiftingLetters(s, shifts)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
