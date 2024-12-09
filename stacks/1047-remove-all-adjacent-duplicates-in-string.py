import unittest

# https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

# python3 -m unittest stacks/1047-remove-all-adjacent-duplicates-in-string.py

class Solution(unittest.TestCase):
    # # Approach: Stack
    # # Time: O(n)
    # # Space: O(n)
    # def removeDuplicates(self, s: str) -> str:
    #     stack = []
    #     for c in s:
    #         if stack and stack[-1] == c:
    #             stack.pop()
    #         else:
    #             stack.append(c)
    #     return "".join(stack)

    # Approach: Two Pointers
    # Time: O(n)
    # Space: O(n)
    def removeDuplicates(self, s: str) -> str:
        res = ['']*len(s)
        j = 0
        for i in range(len(s)):
            res[j] = s[i]
            while j > 0 and res[j-1] == res[j]:
                j -= 2
            j += 1
        return "".join(res[:j])

    def test(self) -> None:
        for s, expected in [
            ("abbaca", "ca"),
            ("azxxzy", "ay"),
            ("aaaaaa", ""),
            ("aaaaaaaaa", "a"),
        ]:
            output = self.removeDuplicates(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
