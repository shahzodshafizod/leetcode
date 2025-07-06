from collections import defaultdict
from math import gcd
from typing import List
import unittest

# https://leetcode.com/problems/count-array-pairs-divisible-by-k/

# python3 -m unittest maths/2183-count-array-pairs-divisible-by-k.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(N^2)
    # # Space: O(1)
    # def countPairs(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     count = 0
    #     for i in range(n):
    #         for j in range(i+1, n):
    #             if nums[i]*nums[j]%k == 0:
    #                 count += 1
    #     return count

    # def countPairs(self, nums: List[int], k: int) -> int:
    #     divisors = []
    #     for num in range(1, k+1):
    #         if k % num == 0:
    #             divisors.append(num)
    #     counter = Counter()
    #     count = 0
    #     for num in nums:
    #         rem = k / gcd(num, k)
    #         count += counter[rem]
    #         for divisor in divisors:
    #             if num % divisor == 0:
    #                 counter[divisor] += 1
    #     return count

    def countPairs(self, nums: List[int], k: int) -> int:
        counter = defaultdict(int)
        count = 0
        for num in nums:
            curr = gcd(num, k)
            for prev, cnt in counter.items():
                if curr * prev % k == 0:
                    count += cnt
            counter[curr] += 1
        return count

    def test(self):
        for nums, k, expected in [
            ([1, 2, 3, 4, 5], 2, 7),
            ([1, 2, 3, 4], 5, 0),
        ]:
            output = self.countPairs(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
