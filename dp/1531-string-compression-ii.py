import unittest

# https://leetcode.com/problems/string-compression-ii/

# python3 -m unittest dp/1531-string-compression-ii.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n*k*26*n) = O(nnk)
    # # Space: O(nnk)
    # def getLengthOfOptimalCompression(self, s: str, k: int) -> int:
    #     n = len(s)
    #     memo = {}
    #     def dp(idx: int, k: int, prev_chr: str, prev_cnt: int) -> int:
    #         if idx+k >= n: return 0
    #         if k < 0: return float("inf")
    #         if (idx,k,prev_chr,prev_cnt) in memo:
    #             return memo[(idx,k,prev_chr,prev_cnt)]
    #         if s[idx] == prev_chr:
    #             incr = 1 if prev_cnt in [1,9,99] else 0
    #             res = incr + dp(idx+1, k, prev_chr, prev_cnt+1)
    #         else:
    #             res = min(
    #                 dp(idx+1, k-1, prev_chr, prev_cnt), # delete s[idx]
    #                 1 + dp(idx+1, k, s[idx], 1), # don't delete
    #             )
    #         memo[(idx,k,prev_chr,prev_cnt)] = res
    #         return res
    #     return dp(0, k, '', 0)

    # Approach #2: Top-Down Dynamic Programming
    # Time: O(nnk)
    # Space: O(nk)
    def getLengthOfOptimalCompression(self, s: str, k: int) -> int:
        # ''->'a', 'a'->'a2', 'a9'->'a10', 'a99'->'a100'
        incr = {1:1,2:1,10:1,100:1}
        n = len(s)
        memo = [[None] * (k+1) for _ in range(n)]
        def dp(idx: int, k: int) -> int:
            if idx+k >= n: return 0
            if k < 0: return 101
            if memo[idx][k] != None:
                return memo[idx][k]
            # delete s[idx]
            res = dp(idx+1, k-1)
            # keep s[idx] and same characters
            same, length, diff = 0, 0, 0
            for j in range(idx, n):
                if s[j] == s[idx]:
                    same += 1
                    length += incr.get(same, 0)
                else:
                    diff += 1
                    if diff > k: break
                res = min(res, length + dp(j+1, k-diff))
            memo[idx][k] = res
            return res
        return dp(0, k)

    def test(self):
        for s, k, expected in [
            ("x", 1, 0),
            ("aaabcccd", 2, 4),
            ("aabbaa", 2, 2),
            ("aaaaaaaaaaa", 0, 3),
            ("llllllllllttttttttt", 1, 4),
            ("eoongjjkjfelnkgkjohfjfjfhkmnmmlinkihhlfipgoejiniol", 13, 32),
            ("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk", 41, 2),
            ("ohoohhoooohhohoooohhhohoohhoohooohhhoohhhooohohhooohhoohhhhhhhooooohhoooohooohooohhohhhhhhohohoohhoo", 74, 3),
            ("earmaypeacebeonearthmerrychristmashohohoooohappyhanukkahhappyholidayshappyholidaysandahappynewyearma", 20, 70),
            ("mkkoqqilqminmrgqprnokiknnrohlqggqniopnpgonjojqihrihqkkhikjgmmggolgmnhoinqqnmlqkoqoihpkkhppoiligjrjnm", 64, 24),
            ("llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll", 25, 3),
        ]:
            output = self.getLengthOfOptimalCompression(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
