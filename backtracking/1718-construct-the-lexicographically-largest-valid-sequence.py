from typing import List
import unittest

# https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/

# python3 -m unittest backtracking/1718-construct-the-lexicographically-largest-valid-sequence.py


class Solution(unittest.TestCase):
    # Approach: Backtracking
    # Time: O(n!)
    # Space: O(n)
    def constructDistancedSequence(self, n: int) -> List[int]:
        size = n * 2 - 1
        seq = [0] * size
        used = [False] * (n + 1)

        def backtrack(idx: int) -> bool:
            if idx == size:
                return True
            if seq[idx] != 0:
                return backtrack(idx + 1)
            for num in range(n, 0, -1):
                if used[num]:
                    continue
                seq[idx] = num
                used[num] = True
                if num == 1:
                    if backtrack(idx + 1):
                        return True
                elif idx + num < size and seq[idx + num] == 0:
                    seq[idx + num] = num
                    if backtrack(idx + 1):
                        return True
                    seq[idx + num] = 0
                seq[idx] = 0
                used[num] = False
            return False

        backtrack(0)
        return seq

    def test(self):
        for n, expected in [
            (3, [3, 1, 2, 3, 2]),
            (5, [5, 3, 1, 4, 3, 5, 2, 4, 2]),
        ]:
            output = self.constructDistancedSequence(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
