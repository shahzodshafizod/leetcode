import unittest

# https://leetcode.com/problems/count-largest-group/

# python3 -m unittest hashes/1399-count-largest-group.py


class Solution(unittest.TestCase):
    def countLargestGroup(self, n: int) -> int:
        count = {}
        size = 0
        for num in range(1, n + 1):
            group = 0
            while num > 0:
                group += num % 10
                num //= 10
            count[group] = count.get(group, 0) + 1
            size = max(size, count[group])
        return sum(1 for cnt in count.values() if cnt == size)

    def test(self):
        for n, expected in [
            (13, 4),
            (2, 2),
        ]:
            output = self.countLargestGroup(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
