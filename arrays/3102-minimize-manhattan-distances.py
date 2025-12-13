from typing import List, Set
import unittest

# https://leetcode.com/problems/minimize-manhattan-distances/

# python3 -m unittest arrays/3102-minimize-manhattan-distances.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # For each point, calculate max Manhattan distance among remaining points
    # # Time: O(n^3) - n iterations, each O(n^2) to find max distance
    # # This approach is too slow for large inputs due to nested loops checking all pairs
    # # Space: O(1)
    # def minimizeMaxDistance(self, points: List[List[int]]) -> int:
    #     def max_manhattan(pts: List[List[int]]) -> int:
    #         if len(pts) <= 1:
    #             return 0
    #         max_dist = 0
    #         for i in range(len(pts)):
    #             for j in range(i + 1, len(pts)):
    #                 dist = abs(pts[i][0] - pts[j][0]) + abs(pts[i][1] - pts[j][1])
    #                 max_dist = max(max_dist, dist)
    #         return max_dist

    #     n = len(points)
    #     min_result = 1 << 31

    #     for skip in range(n):
    #         remaining = [points[i] for i in range(n) if i != skip]
    #         max_dist = max_manhattan(remaining)
    #         min_result = min(min_result, max_dist)

    #     return min_result

    # Approach #2: Optimized - Transform coordinates and use geometry
    # Manhattan distance can be transformed: |x1-x2| + |y1-y2| = max(|(x1+y1)-(x2+y2)|, |(x1-y1)-(x2-y2)|)
    # Transform points to (x+y, x-y), then max Manhattan = max of ranges in both dimensions
    # Time: O(n^2) - try removing each point, O(n) to recalculate max
    # Space: O(n) for transformed coordinates
    def minimizeMaxDistance(self, points: List[List[int]]) -> int:
        n = len(points)

        # Transform coordinates
        u = [x + y for x, y in points]
        v = [x - y for x, y in points]

        def get_max_dist_without(skip_idx: int) -> int:
            # Find max and min in u and v, excluding skip_idx
            u_vals = [u[i] for i in range(n) if i != skip_idx]
            v_vals = [v[i] for i in range(n) if i != skip_idx]

            if not u_vals:
                return 0

            u_range = max(u_vals) - min(u_vals)
            v_range = max(v_vals) - min(v_vals)

            return max(u_range, v_range)

        # Find initial max distance
        u_range = max(u) - min(u)
        v_range = max(v) - min(v)
        current_max = max(u_range, v_range)

        # Try removing points that contribute to max distance
        candidates: Set[int] = set()

        # Find points at extremes
        u_min_idx = u.index(min(u))
        u_max_idx = u.index(max(u))
        v_min_idx = v.index(min(v))
        v_max_idx = v.index(max(v))

        candidates.update([u_min_idx, u_max_idx, v_min_idx, v_max_idx])

        min_result = current_max
        for skip in candidates:
            max_dist = get_max_dist_without(skip)
            min_result = min(min_result, max_dist)

        return min_result

    def test(self):
        for points, expected in [
            ([[3, 10], [5, 15], [10, 2], [4, 4]], 12),
            ([[1, 1], [1, 1], [1, 1]], 0),
            ([[1, 2], [3, 4], [5, 6]], 4),
        ]:
            output = self.minimizeMaxDistance(points)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
