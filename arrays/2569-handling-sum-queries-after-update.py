from typing import List
import unittest

# https://leetcode.com/problems/handling-sum-queries-after-update/

# python3 -m unittest arrays/2569-handling-sum-queries-after-update.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # For each query:
    # # - Type 1: Flip each bit in range [l, r] individually - O(n) per query
    # # - Type 2: Update sum by adding count of 1s * p - O(1)
    # # - Type 3: Return current sum - O(1)
    # #
    # # Time Complexity: O(q * n) where q = number of queries, n = length of nums1
    # # Space Complexity: O(1) excluding output array
    # # This approach will TLE for large inputs (n, q up to 10^5)
    # def handleQuery(self, nums1: List[int], nums2: List[int], queries: List[List[int]]) -> List[int]:
    #     result: List[int] = []
    #     sum2 = sum(nums2)
    #     count1 = sum(nums1)

    #     for query in queries:
    #         if query[0] == 1:
    #             # Flip each bit in range [l, r]
    #             l, r = query[1], query[2]
    #             for i in range(l, r + 1):
    #                 if nums1[i] == 0:
    #                     nums1[i] = 1
    #                     count1 += 1
    #                 else:
    #                     nums1[i] = 0
    #                     count1 -= 1
    #         elif query[0] == 2:
    #             # Add count1 * p to sum2
    #             sum2 += count1 * query[1]
    #         else:
    #             # Return current sum
    #             result.append(sum2)

    #     return result

    # Approach 2: Optimized - Segment Tree with Lazy Propagation
    # Use a segment tree to efficiently handle range flip operations.
    # Each node stores the count of 1s in its range.
    # Lazy propagation defers flip operations until needed.
    #
    # Key insights:
    # 1. For Type 2 queries, we only need count of 1s in nums1, not the array itself
    # 2. Flipping a range [l, r] changes count by: (r - l + 1) - 2 * count_in_range
    # 3. Segment tree handles range flips in O(log n) instead of O(n)
    #
    # Time Complexity: O((n + q) * log n) where q = number of queries
    # Space Complexity: O(n) for segment tree
    def handleQuery(self, nums1: List[int], nums2: List[int], queries: List[List[int]]) -> List[int]:
        # Segment Tree with Lazy Propagation
        class SegmentTree:
            def __init__(self, arr: List[int]):
                self.n = len(arr)
                # Tree stores count of 1s in each segment
                self.tree = [0] * (4 * self.n)
                # Lazy array stores pending flip operations
                self.lazy = [False] * (4 * self.n)
                self._build(arr, 0, 0, self.n - 1)

            def _build(self, arr: List[int], node: int, start: int, end: int):
                """Build the segment tree"""
                if start == end:
                    self.tree[node] = arr[start]
                else:
                    mid = (start + end) // 2
                    left_child = 2 * node + 1
                    right_child = 2 * node + 2
                    self._build(arr, left_child, start, mid)
                    self._build(arr, right_child, mid + 1, end)
                    self.tree[node] = self.tree[left_child] + self.tree[right_child]

            def _push(self, node: int, start: int, end: int):
                """Push lazy updates down to children"""
                if self.lazy[node]:
                    # Flip current node: count of 1s becomes (total - count of 1s)
                    self.tree[node] = (end - start + 1) - self.tree[node]

                    # Propagate to children if not a leaf
                    if start != end:
                        left_child = 2 * node + 1
                        right_child = 2 * node + 2
                        self.lazy[left_child] = not self.lazy[left_child]
                        self.lazy[right_child] = not self.lazy[right_child]

                    self.lazy[node] = False

            def flip_range(self, l: int, r: int):
                """Flip all bits in range [l, r]"""
                self._flip(0, 0, self.n - 1, l, r)

            def _flip(self, node: int, start: int, end: int, l: int, r: int):
                """Recursively flip range [l, r]"""
                # Push pending updates
                self._push(node, start, end)

                # No overlap
                if start > r or end < l:
                    return

                # Complete overlap
                if start >= l and end <= r:
                    self.lazy[node] = True
                    self._push(node, start, end)
                    return

                # Partial overlap
                mid = (start + end) // 2
                left_child = 2 * node + 1
                right_child = 2 * node + 2
                self._flip(left_child, start, mid, l, r)
                self._flip(right_child, mid + 1, end, l, r)

                # Update current node
                self._push(left_child, start, mid)
                self._push(right_child, mid + 1, end)
                self.tree[node] = self.tree[left_child] + self.tree[right_child]

            def get_total_ones(self) -> int:
                """Get total count of 1s in the array"""
                self._push(0, 0, self.n - 1)
                return self.tree[0]

        # Initialize segment tree with nums1
        seg_tree = SegmentTree(nums1)
        sum2 = sum(nums2)
        result: List[int] = []

        for query in queries:
            if query[0] == 1:
                # Flip range [l, r]
                l, r = query[1], query[2]
                seg_tree.flip_range(l, r)
            elif query[0] == 2:
                # Add count of 1s * p to sum2
                count1 = seg_tree.get_total_ones()
                sum2 += count1 * query[1]
            else:
                # Return current sum
                result.append(sum2)

        return result

    def test(self):
        for nums1, nums2, queries, expected in [
            ([1, 0, 1], [0, 0, 0], [[1, 1, 1], [2, 1, 0], [3, 0, 0]], [3]),
            ([1], [5], [[2, 0, 0], [3, 0, 0]], [5]),
            (
                [1, 0, 1, 0, 1],
                [1, 2, 3, 4, 5],
                [[1, 0, 2], [2, 1, 0], [3, 0, 0], [1, 1, 3], [2, 2, 0], [3, 0, 0]],
                [17, 23],
            ),
            (
                [0, 0, 0],
                [1, 2, 3],
                [[1, 0, 2], [2, 3, 0], [3, 0, 0], [1, 1, 1], [2, 2, 0], [3, 0, 0]],
                [15, 19],
            ),
        ]:
            output = self.handleQuery(nums1, nums2, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
