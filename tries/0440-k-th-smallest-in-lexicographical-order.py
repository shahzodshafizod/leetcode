import unittest

# https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/

# python3 -m unittest tries/0440-k-th-smallest-in-lexicographical-order.py

class Solution(unittest.TestCase):
    # Approach: Prefix Tree
    # Time: O((Log N)^2)
    # Space: O(1)
    def findKthNumber(self, n: int, k: int) -> int:
        def count(curr: int) -> int: # Time: O(Log N)
            neighbor = curr+1
            cnt = 0
            while curr <= n:
                cnt += min(n+1, neighbor) - curr
                curr *= 10
                neighbor *= 10
            return cnt
        curr = 1
        k -= 1
        while k > 0: # Time: O(Log N)
            cnt = count(curr)
            if k >= cnt:
                curr += 1
                k -= cnt
            else:
                curr *= 10
                k -= 1
        return curr

    def testFindKthNumber(self) -> None:
        for n, k, expected in [
            (13, 2, 10),
		    (1, 1, 1),
            (1000000000, 1000000000, 999999999),
		    (957747794, 424238336, 481814499),
        ]:
            output = self.findKthNumber(n, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
