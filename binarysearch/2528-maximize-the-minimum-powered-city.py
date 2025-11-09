from typing import List
import unittest

# https://leetcode.com/problems/maximize-the-minimum-powered-city/

# python3 -m unittest binarysearch/2528-maximize-the-minimum-powered-city.py


class Solution(unittest.TestCase):
    # Approach: Binary Search
    # Time: O(n log (s+k)), s=sum(stations)
    # Space: O(n)
    def maxPower(self, stations: List[int], r: int, k: int) -> int:
        n = len(stations)
        line = [0] * (n + 1)
        for i, power in enumerate(stations):
            left = max(i - r, 0)
            right = min(i + r, n - 1)
            line[left] += power
            line[right + 1] -= power

        def maximize(target: int) -> bool:
            presum, quota = 0, k
            diff = line.copy()
            for i in range(n):
                presum += diff[i]
                if presum + quota < target:
                    return False
                if presum < target:
                    delta = target - presum
                    quota -= delta
                    presum = target
                    right = min(i + 2 * r, n - 1)
                    diff[right + 1] -= delta
            return True

        low, high = min(stations), sum(stations) + k
        ans = low
        while low <= high:
            mid = (low + high) // 2
            if maximize(mid):
                ans = mid
                low = mid + 1
            else:
                high = mid - 1

        return ans

    def test(self):
        for stations, r, k, expected in [
            ([1, 2, 4, 5, 0], 1, 2, 5),
            ([4, 4, 4, 4], 0, 3, 4),
            ([2, 10, 12, 3], 0, 14, 9),
        ]:
            output = self.maxPower(stations, r, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
