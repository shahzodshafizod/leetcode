import unittest

# https://leetcode.com/problems/maximize-the-number-of-partitions-after-operations/

# python3 -m unittest dp/3003-maximize-the-number-of-partitions-after-operations.py


class Solution(unittest.TestCase):
    def maxPartitionsAfterOperations(self, s: str, k: int) -> int:
        memo = {}
        def dp(i: int, mask: int, changed: bool) -> int:
            if i == len(s): return 1
            key = (i, mask, changed)
            if key in memo: return memo[key]
            best = 0
            cur = ord(s[i]) - ord('a')
            nmask = mask | (1 << cur)
            if nmask.bit_count() <= k:
                best = max(best, dp(i+1, nmask, changed))
            else:
                best = max(best, 1 + dp(i+1, 1 << cur, changed))
            if not changed:
                for nxt in range(26):
                    if nxt == cur: continue
                    nmask = mask | (1 << nxt)
                    if nmask.bit_count() <= k:
                        best = max(best, dp(i+1, nmask, True))
                    else:
                        best = max(best, 1 + dp(i+1, 1 << nxt, True))
            memo[key] = best
            return best
        return dp(0, 0, False)

    def test(self):
        for s, k, expected in [
            ("accca", 2, 3),
            ("aabaab", 3, 1),
            ("xxyz", 1, 4),
        ]:
            output = self.maxPartitionsAfterOperations(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
