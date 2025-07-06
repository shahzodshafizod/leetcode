import unittest

# https://leetcode.com/problems/first-bad-version/

# python3 -m unittest binarysearch/0278-first-bad-version.py


# The isBadVersion API is already defined for you.
def isBadVersion(version: int) -> bool:
    return version >= bad


class Solution(unittest.TestCase):
    def firstBadVersion(self, n: int) -> int:
        left, right = 1, n
        while left < right:
            mid = (left + right) // 2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid + 1
        return right

    def testFirstBadVersion(self) -> None:
        global bad  # pylint: disable=global-variable-undefined
        for n, local_bad, expected in [
            (5, 4, 4),
            (1, 1, 1),
            (3, 2, 2),
        ]:
            bad = local_bad
            output = self.firstBadVersion(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
