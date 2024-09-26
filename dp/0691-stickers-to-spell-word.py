from collections import Counter, defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/stickers-to-spell-word/

# python3 -m unittest dp/0691-stickers-to-spell-word.py

class Solution(unittest.TestCase):
    # Approach: Top-Down Dynamic Programming
    # N=len(stickers)
    # M=len(target)
    # for every letter in target we decide whether to choose the sticker or not
    # in the worst case for every letter in target we choose every sticker
    # Time: O(2^N ∗M)
    # Space: O(2^N ∗M)
    def minStickers(self, stickers: List[str], target: str) -> int:
        MAX_INT = 16
        stickers = [Counter(sticker) for sticker in stickers]
        memo = defaultdict(int)
        def dfs(target: str) -> int:
            if not target:
                return 0
            if target in memo:
                return memo[target]
            target_map = Counter(target)
            count = MAX_INT
            for sticker in stickers:
                if target[0] not in sticker:
                    continue
                next = target_map - sticker & target_map
                next = "".join([ch*cnt for ch,cnt in next.items()])
                count = min(count, 1 + dfs(next)) # +1 mean we chose this sticker
            memo[target] = count
            return count
        count = dfs(target)
        return count if count != MAX_INT else -1

    def testMinStickers(self) -> None:
        for stickers, target, expected in [
            (["with","example","science"], "thehat", 3),
            (["notice","possible"], "basicbasic", -1),
            (["charge","mind","bottom"], "centorder", 4),
            (["fly","me","charge","mind","bottom"], "centorder", 4),
            (["these", "guess", "about", "garden", "him"], "atomher", 3),
        ]:
            output = self.minStickers(stickers, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
