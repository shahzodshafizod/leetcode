from typing import List
import unittest

# https://leetcode.com/problems/lexicographical-numbers/

# python3 -m unittest strings/0386-lexicographical-numbers.py

class Solution(unittest.TestCase):
    # # Approach #1: Python One Liner
    # # Time: O(nlogn)
    # # Space: O(n)
    # def lexicalOrder(self, n: int) -> List[int]:
    #     return list(map(int,sorted(map(str,range(1,n+1)))))

    # # Approach #2: Depth-First Search
    # # Time: O(n)
    # # Space: O(log(10)n), for recursion stack
    # def lexicalOrder(self, n: int) -> List[int]:
    #     nums = []
    #     def dfs(num: int):
    #         if num > n: return
    #         nums.append(num)
    #         for k in range(10):
    #             dfs(num*10+k)
    #     for num in range(1, 10):
    #         dfs(num)
    #     return nums

    # Approach #3: Iterative
    # Time: O(n)
    # Space: O(1)
    def lexicalOrder(self, n: int) -> List[int]:
        nums = [None] * n
        num = 1
        for idx in range(n):
            nums[idx] = num
            if num*10 <= n:
                num *= 10
            else:
                while num == n or num%10 == 9:
                    num //= 10
                num += 1
        return nums

    def test(self):
        for n, expected in [
            (13, [1,10,11,12,13,2,3,4,5,6,7,8,9]),
            (2, [1,2]),
        ]:
            output = self.lexicalOrder(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
