import unittest

# https://leetcode.com/problems/rotate-string/

# python3 -m unittest strings/0796-rotate-string.py

class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nn)
    # # Space: O(1)
    # def rotateString(self, s: str, goal: str) -> bool:
    #     if len(s) != len(goal): return False
    #     n = len(s)
    #     for start in range(n):
    #         sid, gid = start, 0
    #         while gid < n:
    #             if s[sid] != goal[gid]: break
    #             gid += 1
    #             sid = (sid + 1) % n
    #         if gid == n: return True
    #     return False

    # # Approach #2: Concatenation
    # # Time: O(n)
    # # Space: O(n)
    # def rotateString(self, s: str, goal: str) -> bool:
    #     # return len(s) == len(goal) and (s+s).find(goal) != -1
    #     return len(s) == len(goal) and s in goal+goal

    # Approach: KMP (Knuth-Morris-Pratt) Algorithm
    # Time: O(n)
    # Space: O(n)
    def rotateString(self, s: str, goal: str) -> bool:
        if len(s) != len(goal): return False
        # 1. LPS setup: O(2xn)
        n = len(goal)
        lps = [0] * n
        prev_lps = 0
        for gid in range(1, n):
            while prev_lps > 0 and goal[prev_lps] != goal[gid]:
                prev_lps = lps[prev_lps-1]
            if goal[prev_lps] == goal[gid]:
                prev_lps += 1
            lps[gid] = prev_lps
        # 2. KMP: O(2xn)
        gid = 0
        for c in s+s:
            while gid > 0 and goal[gid] != c:
                gid = lps[gid-1]
            if goal[gid] == c: gid += 1
            if gid == n: return True
        return False

    def test(self):
        for s, goal, expected in [
            ("a", "a", True),
            ("bcad", "cdba", False),
            ("bcad", "abcd", False),
            ("abcd", "cdeab", False),
            ("abcde", "abce", False),
            ("abcde", "cdeab", True),
		    ("abcde", "abced", False),
            ("aaaaaaaaaa", "aaaaaaaaaaaa", False),
            ("bbbacddceeb", "ceebbbbacdd", True),
            ("bbbacddceeb", "ceebbbbacdd", True),
            ("abcdebcdefin", "abcdebcdefin", True),
            ("aaaaaaaaaaaaaaabc", "caaaaaaaaaaaaaaab", True),
            ("dvadererrerdvaddf", "rerdvaddfdvaderer", True),
        ]:
            output = self.rotateString(s, goal)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
