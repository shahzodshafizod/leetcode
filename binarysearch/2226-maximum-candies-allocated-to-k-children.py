from typing import List
import unittest

# https://leetcode.com/problems/maximum-candies-allocated-to-k-children/

# python3 -m unittest binarysearch/2226-maximum-candies-allocated-to-k-children.py


class Solution(unittest.TestCase):
    def maximumCandies(self, candies: List[int], k: int) -> int:
        def can_allocate(target: int) -> bool:
            children = 0
            for candy in candies:
                children += candy // target
                if children >= k:
                    return True
            return False

        total = sum(candies)
        if total < k:
            return 0
        low, high = 1, total // k
        count = 0
        while low <= high:
            mid = (low + high) // 2
            if can_allocate(mid):
                count = mid
                low = mid + 1
            else:
                high = mid - 1
        return count

    def test(self):
        for candies, k, expected in [
            ([5, 8, 6], 3, 5),
            ([2, 5], 11, 0),
            ([4, 7, 5], 4, 3),
            ([1, 2, 3, 4, 10], 5, 3),
            ([5, 6, 4, 10, 10, 1, 1, 2, 2, 2], 9, 3),
            ([1, 2, 6, 8, 6, 7, 3, 5, 2, 5], 3, 6),
            ([4, 9, 4, 7, 8, 10, 3, 10, 3, 9], 9, 4),
            ([9, 10, 1, 2, 10, 9, 9, 10, 2, 2], 3, 10),
        ]:
            output = self.maximumCandies(candies, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
