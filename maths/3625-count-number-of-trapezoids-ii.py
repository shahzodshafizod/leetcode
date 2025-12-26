from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/count-number-of-trapezoids-ii/

# python3 -m unittest maths/3625-count-number-of-trapezoids-ii.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force - Check All Quadrilaterals
    # # Try all combinations of 4 points and check if they form a trapezoid.
    # # A trapezoid has at least one pair of parallel sides, but NOT all collinear.
    # # Key: For any 4 points, there are exactly 3 ways to pair them into
    # # opposite sides: (p1,p2)-(p3,p4), (p1,p3)-(p2,p4), (p1,p4)-(p2,p3)
    # # If exactly one or two pairs have the same slope (not all three), it's a trapezoid.
    # #
    # # Time Complexity: O(n^4) where n = number of points
    # # Space Complexity: O(1)
    # # This approach will TLE for n = 500 (over 6 billion operations)
    # def countTrapezoids(self, points: List[List[int]]) -> int:
    #     def get_slope(p1: List[int], p2: List[int]) -> tuple[int, int]:
    #         """Calculate slope between two points, return (dy, dx) in reduced form"""
    #         dx = p2[0] - p1[0]
    #         dy = p2[1] - p1[1]
    #         if dx == 0:
    #             return (1, 0)  # Vertical line
    #         if dy == 0:
    #             return (0, 1)  # Horizontal line
    #         # Reduce fraction and normalize sign
    #         g = gcd(abs(dx), abs(dy))
    #         dx, dy = dx // g, dy // g
    #         # Normalize: ensure dx is positive
    #         if dx < 0:
    #             dx, dy = -dx, -dy
    #         return (dy, dx)

    #     def is_trapezoid(p1: List[int], p2: List[int], p3: List[int], p4: List[int]) -> bool:
    #         """Check if 4 points form a trapezoid (at least one pair of parallel segments)"""
    #         # For 4 points, check all 3 ways to pair them as opposite sides
    #         # Check if any pair has parallel segments (same slope)
    #         slope_12_34 = get_slope(p1, p2) == get_slope(p3, p4)
    #         slope_13_24 = get_slope(p1, p3) == get_slope(p2, p4)
    #         slope_14_23 = get_slope(p1, p4) == get_slope(p2, p3)

    #         # Count how many pairs are parallel
    #         parallel_count = sum([slope_12_34, slope_13_24, slope_14_23])

    #         # Trapezoid: at least one pair parallel, but NOT all collinear
    #         # If all 3 pairs are parallel, points are collinear (not a valid quadrilateral)
    #         return parallel_count > 0 and parallel_count < 3

    #     n = len(points)
    #     count = 0

    #     # Try all combinations of 4 points
    #     for i in range(n):
    #         for j in range(i + 1, n):
    #             for k in range(j + 1, n):
    #                 for l in range(k + 1, n):
    #                     if is_trapezoid(points[i], points[j], points[k], points[l]):
    #                         count += 1

    #     return count

    # Approach 2: Optimized - Hash by Slope and Intercept
    # Key insights from hints:
    # 1. Hash every point-pair by its slope (normalized with GCD)
    # 2. For each slope, group segments by their line (using intercept)
    # 3. Count pairs of segments with same slope but different lines
    # 4. Subtract overcounts from parallelograms (counted twice)
    #
    # Algorithm:
    # - For each pair of points (segment), calculate:
    #   - Slope (as floating point for simplicity)
    #   - Intercept (to distinguish parallel lines)
    #   - Midpoint (to identify parallelograms)
    # - Count all pairs of parallel segments (potential trapezoids)
    # - Subtract segments sharing endpoints
    # - Subtract parallelogram overcounts
    #
    # Time Complexity: O(n^2) for processing pairs + O(n^2) for counting
    # Space Complexity: O(n^2) for storing segment information
    def countTrapezoids(self, points: List[List[int]]) -> int:
        n = len(points)
        inf = int(1e9 + 7)
        slope_to_intercept: dict[float, list[float]] = defaultdict(list)
        mid_to_slope: dict[int, list[float]] = defaultdict(list)
        ans = 0

        # Process all pairs of points (segments)
        for i in range(n):
            x1, y1 = points[i]
            for j in range(i + 1, n):
                x2, y2 = points[j]
                dx = x1 - x2
                dy = y1 - y2

                # Calculate slope
                if x2 == x1:
                    # Vertical line: use special marker
                    k = float(inf)
                    b = float(x1)  # Intercept is x-coordinate
                else:
                    # Slope = (y2 - y1) / (x2 - x1)
                    k = (y2 - y1) / (x2 - x1)
                    # Intercept: y - slope * x
                    # b = y1 - k * x1 = (y1 * dx - x1 * dy) / dx
                    b = (y1 * dx - x1 * dy) / dx

                # Midpoint of segment (to identify parallelograms)
                # Use encoding to avoid tuple hashing issues
                mid = (x1 + x2) * 10000 + (y1 + y2)

                slope_to_intercept[k].append(b)
                mid_to_slope[mid].append(k)

        # Count all pairs of parallel segments (trapezoid candidates)
        for intercepts in slope_to_intercept.values():
            if len(intercepts) < 2:
                continue

            # Group by intercept (segments on same line don't form trapezoid)
            intercept_count: dict[float, int] = defaultdict(int)
            for b_val in intercepts:
                intercept_count[b_val] += 1

            # Count pairs of segments on different parallel lines
            # For each new intercept, count pairs with all previous ones
            total_sum = 0
            for count in intercept_count.values():
                ans += total_sum * count
                total_sum += count

        # Subtract parallelogram overcounts
        # Parallelograms are counted twice (once for each pair of parallel sides)
        for slopes in mid_to_slope.values():
            if len(slopes) < 2:
                continue

            # Group by slope
            slope_count: dict[float, int] = defaultdict(int)
            for k_val in slopes:
                slope_count[k_val] += 1

            # Subtract pairs of segments with same slope and same midpoint
            total_sum = 0
            for count in slope_count.values():
                ans -= total_sum * count
                total_sum += count

        return ans

    def test(self):
        for points, expected in [
            ([[-3, 2], [3, 0], [2, 3], [3, 2], [2, -3]], 2),
            ([[0, 0], [1, 0], [0, 1], [2, 1]], 1),
            ([[1, 1], [2, 2], [3, 3], [4, 4]], 0),  # All collinear
            ([[0, 0], [1, 0], [0, 1], [1, 1]], 1),  # Square
        ]:
            output = self.countTrapezoids(points)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
