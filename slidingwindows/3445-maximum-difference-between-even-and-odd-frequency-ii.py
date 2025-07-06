import unittest

# even - even = even
# even - odd = odd
# odd - odd = even
# odd - even = odd

# [e]e - [o]e = oe
# [e]o - [o]o = oe
# [o]e - [e]e = oe
# [o]o - [e]o = oe
# it means the second remains the same, but the first changes

# https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-ii/

# python3 -m unittest slidingwindows/3445-maximum-difference-between-even-and-odd-frequency-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Sliding Window
    # # Time: O(nn)
    # # Space: O(1)
    # def maxDifference(self, s: str, k: int) -> int:
    #     n = len(s)
    #     diff = -n
    #     for k in range(k,n+1):
    #         count = [0] * 5
    #         for idx in range(k-1):
    #             count[int(s[idx])] += 1
    #         for idx in range(k-1,n):
    #             count[int(s[idx])] += 1
    #             odd, even = 0, n
    #             for cnt in count:
    #                 if cnt&1 == 1: odd = max(cnt, odd)
    #                 elif cnt > 0: even = min(cnt, even)
    #             diff = max(diff, odd - even)
    #             count[int(s[idx-k+1])] -= 1
    #     return diff

    # Approach #2: Prefix Sum + Sliding Window
    # Time: O(n + mm), m=5
    # Space: O(n)
    def maxDifference(self, s: str, k: int) -> int:
        def get_stat(cnt_a: int, cnt_b: int) -> int:
            return ((cnt_a & 1) << 1) + (cnt_b & 1)

        n = len(s)

        def calc_difference(a: str, b: str) -> int:
            best = [0, n, n, n] * 4
            best_b = [0, None, None, None] * 4
            cnt_a, cnt_b = 0, 0
            prev_a, prev_b = 0, 0
            max_diff = -n
            for idx in range(n):
                cnt_a += s[idx] == a
                cnt_b += s[idx] == b
                if idx + 1 >= k:
                    stat = get_stat(cnt_a, cnt_b) ^ 0b10
                    if best[stat] < n and cnt_b > best_b[stat]:
                        # B must have frequency > 0 in the subarray
                        max_diff = max(max_diff, cnt_a - cnt_b - best[stat])
                    # slide window to right
                    prev_a += s[idx + 1 - k] == a
                    prev_b += s[idx + 1 - k] == b
                    prev_stat = get_stat(prev_a, prev_b)
                    prev_diff = prev_a - prev_b
                    if prev_diff < best[prev_stat]:
                        best[prev_stat] = prev_diff
                        best_b[prev_stat] = prev_b
            return max_diff

        vals = ['0', '1', '2', '3', '4']
        return max(calc_difference(a, b) for a in vals for b in vals if a != b)

    def test(self):
        for s, k, expected in [
            ("12233", 4, -1),
            ("1122211", 3, 1),
            ("110", 3, -1),
            ("2222130", 2, -1),
            ("11131340", 8, -1),
            ("111112222233333", 5, 3),
        ]:
            output = self.maxDifference(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
