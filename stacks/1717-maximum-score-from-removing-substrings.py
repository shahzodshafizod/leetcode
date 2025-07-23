import unittest

# https://leetcode.com/problems/maximum-score-from-removing-substrings/

# python3 -m unittest stacks/1717-maximum-score-from-removing-substrings.py


class Solution(unittest.TestCase):
    # # Approach #1: Stack
    # # Time: O(2n) = O(n)
    # # Space: O(n)
    # def maximumGain(self, s: str, x: int, y: int) -> int:
    #     a, b = 'a', 'b'
    #     if x < y:
    #         x, y = y, x
    #         a, b = b, a
    #     bayts, size = list(s), len(s)

    #     def get_score(a: str, b: str, score: int) -> int:
    #         nonlocal bayts, size
    #         res, ptr = 0, 0
    #         for i in range(size):
    #             if ptr > 0 and bayts[ptr - 1] == a and bayts[i] == b:
    #                 res += score
    #                 ptr -= 1
    #             else:
    #                 bayts[ptr] = bayts[i]
    #                 ptr += 1
    #         size = ptr
    #         return res

    #     return get_score(a, b, x) + get_score(b, a, y)

    # Approach #2: Greedy (Counting)
    # Time: O(n)
    # Space: O(1)
    def maximumGain(self, s: str, x: int, y: int) -> int:
        a, b = 'a', 'b'
        if x < y:
            x, y = y, x
            a, b = b, a

        score, acnt, bcnt = 0, 0, 0
        for c in s:
            if c == a:
                acnt += 1
            elif c == b:
                if acnt > 0:
                    score += x
                    acnt -= 1
                else:
                    bcnt += 1
            else:
                score += min(acnt, bcnt) * y
                acnt, bcnt = 0, 0
        score += min(acnt, bcnt) * y
        return score

    def test(self):
        for s, x, y, expected in [
            ("cdbcbbaaabab", 4, 5, 19),
            ("aabbaaxybbaabb", 5, 4, 20),
            ("abbaab", 1, 100, 201),
        ]:
            output = self.maximumGain(s, x, y)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
