from typing import List, Set
import unittest

# https://leetcode.com/problems/minimum-number-of-people-to-teach/

# python3 -m unittest greedy/1733-minimum-number-of-people-to-teach.py


class Solution(unittest.TestCase):
    def minimumTeachings(self, n: int, languages: List[List[int]], friendships: List[List[int]]) -> int:
        slanguages: List[Set[int]] = [set()] + [set(langs) for langs in languages]
        unconnected: Set[int] = set()
        for f1, f2 in friendships:
            if slanguages[f1] & slanguages[f2]:
                continue  # & is intersection of sets
            unconnected.add(f1)
            unconnected.add(f2)
        freq: List[int] = [0] * (n + 1)
        for friend in unconnected:
            for lang in languages[friend - 1]:
                freq[lang] += 1
        return len(unconnected) - max(freq)

    def test(self):
        for n, languages, friendships, expected in [
            (2, [[1], [2], [1, 2]], [[1, 2], [1, 3], [2, 3]], 1),
            (3, [[2], [1, 3], [1, 2], [3]], [[1, 4], [1, 2], [3, 4], [2, 3]], 2),
        ]:
            output = self.minimumTeachings(n, languages, friendships)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
