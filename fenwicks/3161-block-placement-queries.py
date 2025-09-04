from typing import List, Dict
import unittest

# https://leetcode.com/problems/block-placement-queries/

# python3 -m unittest fenwicks/3161-block-placement-queries.py


class FenwickTreeMax:
    def __init__(self, n: int):
        self.n = n
        self.tree = [0] * (n + 1)

    def update(self, idx: int, val: int):
        while idx <= self.n:
            self.tree[idx] = max(self.tree[idx], val)
            idx += idx & -idx

    def query(self, idx: int) -> int:
        ans = 0
        while idx > 0:
            ans = max(ans, self.tree[idx])
            idx -= idx & -idx
        return ans


class Solution(unittest.TestCase):
    def getResults(self, queries: List[List[int]]) -> List[bool]:
        # Collect all obstacle positions (type 1 queries)
        blocks = sorted([0, 50001] + [x for t, x, *_ in queries if t == 1])
        n = len(blocks)

        # Initialize map with neighbors for each obstacle/block
        ptrs: Dict[int, List[int]] = {}
        ptrs[blocks[0]] = [0, blocks[1]]
        for i in range(1, n - 1):
            ptrs[blocks[i]] = [blocks[i - 1], blocks[i + 1]]
        ptrs[blocks[-1]] = [blocks[-2], 0]

        # Fenwick tree for max interval length queries and updates
        fenw = FenwickTreeMax(n - 1)
        for i in range(n - 1):
            length = blocks[i + 1] - blocks[i]
            fenw.update(i + 1, length)  # i+1 for 1-based index

        # Map block positions to Fenwick tree indices
        indices = {pos: idx + 1 for idx, pos in enumerate(blocks)}  # 1-based indexing for Fenwicks

        ans: List[bool] = []
        # We process queries from last to first
        for t, x, *sz in reversed(queries):
            if t == 1:
                # Remove obstacle x, merge intervals around it
                prv, nxt = ptrs[x]

                # Remove old intervals from fenwick, effectively by resetting to zero
                # Fenwicks handle max updates, so we do update with 0 to remove effect
                fenw.update(indices[prv], 0)
                fenw.update(indices[x], 0)

                # Merge intervals in ptrs structure
                ptrs[prv][1] = nxt
                ptrs[nxt][0] = prv

                # Insert merged interval length
                fenw.update(indices[prv], nxt - prv)

            else:
                y = sz[0]

                # Need to find if there's an interval whose length >= y and start + y <= x
                # Traverse intervals in fenw and check condition via binary search on blocks
                # We'll iterate over intervals that have length >= y using fenw max query approach

                # We can check max interval length overall from fenw root:
                if y > x or x > fenw.query(n - 1):
                    ans.append(False)
                    continue

                # Since fenw queries max prefix, we try a binary search over blocks to find a suitable interval
                left, right = 1, n - 1
                while left <= right:
                    mid = (left + right) // 2
                    if fenw.query(mid) >= y:
                        right = mid - 1
                    else:
                        left = mid + 1

                # Now left points to the smallest index with max interval >= y
                # Check if interval start + y <= x
                # interval start is blocks[left-1]
                ans.append(left <= n - 1 and blocks[left - 1] + y <= x)

        ans.reverse()
        return ans

    def test(self):
        for queries, expected in [
            ([[1, 2], [2, 3, 3], [2, 3, 1], [2, 2, 2]], [False, True, True]),
            ([[1, 7], [2, 7, 6], [1, 2], [2, 7, 5], [2, 7, 6]], [True, True, False]),
        ]:
            output = self.getResults(queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
