import unittest

# https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/

# python3 -m unittest greedy/1415-the-k-th-lexicographical-string-of-all-happy-strings-of-length-n.py

class Solution(unittest.TestCase):
    def getHappyString(self, n: int, k: int) -> str:
        happy = [''] * n
        partition_len = 1 << n
        for idx in range(n):
            partition_len //= 2
            i = 0
            for c in "abc":
                if happy[idx-1] == c:
                    continue
                if k <= (i+1) * partition_len:
                    happy[idx] = c
                    k -= i * partition_len
                    break
                i += 1
        return "".join(happy)

    def test(self):
        for n, k, expected in [
            (1, 3, "c"),
            (1, 4, ""),
            (3, 9, "cab"),
        ]:
            output = self.getHappyString(n, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
