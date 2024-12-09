from collections import deque
import unittest

# https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/

# python3 -m unittest dp/1553-minimum-number-of-days-to-eat-n-oranges.py

class Solution(unittest.TestCase):
    # # Approach: Breadth-First Search
    # # Time: O(logn)
    # # Space: O(n)
    # def minDays(self, n: int) -> int:
    #     seen = {}
    #     queue = deque([n])
    #     seen = set([n])
    #     days = 0
    #     while queue:
    #         for _ in range(len(queue)):
    #             n = queue.popleft()
    #             if n == 0: return days
    #             if n%3 == 0 and n//3 not in seen:
    #                 queue.append(n//3)
    #                 seen.add(n//3)
    #             if n%2 == 0 and n//2 not in seen:
    #                 queue.append(n//2)
    #                 seen.add(n//2)
    #             if n-1 not in seen:
    #                 queue.append(n-1)
    #                 seen.add(n-1)
    #         days += 1
    #     return -1

    # Approach: Top-Down Dynamic Programming
    # Time: O(logn)
    # Space: O(logn)
    def minDays(self, n: int) -> int:
        # We have three options:
        # - east 1, remains n-1: 1+n-1 = n
        # - east n/2, remains n/2, n/2+n/2 = n
        # - east 2*n/3, remains n/3, 2*n/3+n/3 = 3*n/3 = n
        memo = {}
        def dp(n: int) -> int:
            if n <= 1: return n
            if n in memo: return memo[n]
            memo[n] = 1 + min(
                n%3 + dp(n//3), # make n divisable to 3 and add the remaining
                n%2 + dp(n//2), # make n divisable to 2 and add the remaining
            )
            return memo[n]
        return dp(n)

    def test(self) -> None:
        for n, expected in [
		    (6, 3),
            (10, 4),
            # (1073741826, 32),
            # (2000000000, 32),
        ]:
            output = self.minDays(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
