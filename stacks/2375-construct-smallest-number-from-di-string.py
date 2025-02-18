import unittest

# https://leetcode.com/problems/construct-smallest-number-from-di-string/

# python3 -m unittest stacks/2375-construct-smallest-number-from-di-string.py

class Solution(unittest.TestCase):
    # # Approach #1: Backtracking
    # # Time: O()
    # # Space: O(n)
    # def smallestNumber(self, pattern: str) -> str:
    #     n = len(pattern)
    #     nums = [0] * (n+1)
    #     used = [False] * 10
    #     def backtrack(idx: int) -> bool:
    #         if idx == n:
    #             return True
    #         for num in range(1, 10):
    #             if used[num]:
    #                 continue
    #             if pattern[idx] == 'I' and nums[idx] > num:
    #                 continue
    #             if pattern[idx] == 'D' and nums[idx] < num:
    #                 continue
    #             nums[idx+1] = num
    #             used[num] = True
    #             if backtrack(idx+1):
    #                 return True
    #             used[num] = False
    #         return False
    #     for num in range(1, 10):
    #         nums[0] = num
    #         used[num] = True
    #         if backtrack(0):
    #             break
    #         used[num] = False
    #     return "".join(str(num) for num in nums)
    
    # Approach #1: Backtracking
    # Time: O()
    # Space: O(n)
    def smallestNumber(self, pattern: str) -> str:
        n = len(pattern)
        stack, num = [], []
        for idx in range(n+1):
            stack.append(idx+1)
            if idx == n or pattern[idx] == 'I':
                while stack:
                    num.append(str(stack.pop()))
        return "".join(num)

    def test(self):
        for pattern, expected in [
            ("IIIDIDDD", "123549876"),
            ("DDD", "4321"),
        ]:
            output = self.smallestNumber(pattern)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
