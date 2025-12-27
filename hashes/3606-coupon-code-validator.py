import unittest
from typing import List, Tuple

# https://leetcode.com/problems/coupon-code-validator/

# python3 -m unittest hashes/3606-coupon-code-validator.py


class Solution(unittest.TestCase):
    # Approach: Filtering + Sorting with Custom Comparator
    # Time: O(n log n) - sorting valid coupons
    # Space: O(n) - storing valid coupons
    def validateCoupons(self, code: List[str], businessLine: List[str], isActive: List[bool]) -> List[str]:
        # Define business line priority order
        priority = {"electronics": 0, "grocery": 1, "pharmacy": 2, "restaurant": 3}
        valid_businesses = set(priority.keys())

        def is_valid_code(c: str) -> bool:
            """Check if code is non-empty and contains only alphanumeric chars and underscores"""
            if not c:
                return False
            return all(ch.isalnum() or ch == '_' for ch in c)

        # Collect valid coupons as (priority, code) tuples
        valid_coupons: List[Tuple[int, str]] = []
        for i in range(len(code)):
            if (isActive[i] and
                is_valid_code(code[i]) and
                businessLine[i] in valid_businesses):
                valid_coupons.append((priority[businessLine[i]], code[i]))

        # Sort by priority first, then by code lexicographically
        valid_coupons.sort()

        # Extract and return just the codes
        return [c for _, c in valid_coupons]

    def test(self):
        for code, businessLine, isActive, expected in [
            (
                ["SAVE20", "", "PHARMA5", "SAVE@20"],
                ["restaurant", "grocery", "pharmacy", "restaurant"],
                [True, True, True, True],
                ["PHARMA5", "SAVE20"],
            ),
            (
                ["GROCERY15", "ELECTRONICS_50", "DISCOUNT10"],
                ["grocery", "electronics", "invalid"],
                [False, True, True],
                ["ELECTRONICS_50"],
            ),
            (
                ["E1", "E2", "G1", "P1", "R1"],
                ["electronics", "electronics", "grocery", "pharmacy", "restaurant"],
                [True, True, True, True, True],
                ["E1", "E2", "G1", "P1", "R1"],
            ),
        ]:
            output = self.validateCoupons(code, businessLine, isActive)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
