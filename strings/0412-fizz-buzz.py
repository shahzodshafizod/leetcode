from typing import List
import unittest

# https://leetcode.com/problems/fizz-buzz/

# python3 -m unittest strings/0412-fizz-buzz.py

class Solution(unittest.TestCase):
    # # Approach #1
    # # Time: O(N)
    # # Space: O(1)
    # def fizzBuzz(self, n: int) -> List[str]:
    #     # return ['Fizz' * (not i % 3) + 'Buzz' * (not i % 5) or str(i) for i in range(1, n+1)]
    #     answer = [""] * n
    #     for idx in range(1, n+1):
    #         if idx%3 == 0:
    #             answer[idx-1] = "Fizz"
    #         if idx%5 == 0:
    #             answer[idx-1] += "Buzz"
    #         if not answer[idx-1]:
    #             answer[idx-1] = str(idx)
    #     return answer

    # # Approach #2
    # # Time: O(N)
    # # Space: O(1)
    # def fizzBuzz(self, n: int) -> List[str]:
    #     source = ["", "Fizz", "Buzz", "FizzBuzz"]
    #     answer = [""] * n
    #     for idx in range(1, n+1):
    #         sid = int(idx%3 == 0) + int(idx%5 == 0) * 2
    #         if sid == 0:
    #             answer[idx-1] = str(idx)
    #         else:
    #             answer[idx-1] = source[sid]
    #     return answer

    # Approach #3. the use of a modulus operator may have "some" impact on time complexity
    # Time: O(N)
    # Space: O(1)
    def fizzBuzz(self, n: int) -> List[str]:
        fizz, buzz = 3, 5
        answer = [""] * n
        for idx in range(1, n+1):
            if idx == fizz:
                answer[idx-1] = "Fizz"
                fizz += 3
            if idx == buzz:
                answer[idx-1] += "Buzz"
                buzz += 5
            if not answer[idx-1]:
                answer[idx-1] = str(idx)
        return answer

    def test(self):
        for n, expected in [
            (3, ["1","2","Fizz"]),
            (5, ["1","2","Fizz","4","Buzz"]),
            (15, ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]),
        ]:
            output = self.fizzBuzz(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
