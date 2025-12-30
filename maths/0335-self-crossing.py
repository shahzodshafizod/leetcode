from typing import List
import unittest

# https://leetcode.com/problems/self-crossing/

# python3 -m unittest maths/0335-self-crossing.py


class Solution(unittest.TestCase):
    # Approach: Pattern Recognition and Geometry
    # Key insight: Due to the counter-clockwise spiral pattern, the i-th line can only
    # cross with lines that are 3, 4, or 5 steps before it.
    #
    # Three crossing patterns:
    # 1. Fourth line crosses first line (i crosses i-3):
    #    Happens when: distance[i] >= distance[i-2] and distance[i-1] <= distance[i-3]
    #
    # 2. Fifth line crosses second line (i crosses i-4):
    #    Happens when: distance[i-1] == distance[i-3] and
    #                  distance[i] + distance[i-4] >= distance[i-2]
    #
    # 3. Sixth line crosses third line (i crosses i-5):
    #    Happens when: distance[i-2] >= distance[i-4] and
    #                  distance[i-3] >= distance[i-1] and
    #                  distance[i-1] + distance[i-5] >= distance[i-3] and
    #                  distance[i] + distance[i-4] >= distance[i-2]
    #
    # Visual patterns:
    # Pattern 1 (4th crosses 1st):    Pattern 2 (5th crosses 2nd):    Pattern 3 (6th crosses 3rd):
    #       0                              0──1                             0──1
    #       │                              │  │                             │  │
    #   3───┼──1                       4───┼──┼──2                      5───┼──2
    #       │  │                           │  │                         │   │  │
    #       └──2                           └──3                         └───┼──3
    #                                                                       │
    #                                                                       4
    #
    # Time: O(n) where n is the length of distance array
    # Space: O(1) - only using a few variables
    def isSelfCrossing(self, distance: List[int]) -> bool:
        n = len(distance)
        if n < 4:
            return False

        for i in range(3, n):
            # Pattern 1: Fourth line crosses first line (i crosses i-3)
            if distance[i] >= distance[i - 2] and distance[i - 1] <= distance[i - 3]:
                return True

            # Pattern 2: Fifth line crosses second line (i crosses i-4)
            if i >= 4:
                if (
                    distance[i - 1] == distance[i - 3]
                    and distance[i] + distance[i - 4] >= distance[i - 2]
                ):
                    return True

            # Pattern 3: Sixth line crosses third line (i crosses i-5)
            if i >= 5:
                if (
                    distance[i - 2] >= distance[i - 4]
                    and distance[i - 3] >= distance[i - 1]
                    and distance[i - 1] + distance[i - 5] >= distance[i - 3]
                    and distance[i] + distance[i - 4] >= distance[i - 2]
                ):
                    return True

        return False

    def test(self):
        for distance, expected in [
            ([2, 1, 1, 2], True),
            ([1, 2, 3, 4], False),
            ([1, 1, 1, 2, 1], True),
            ([1, 1, 2, 1, 1], True),
            ([3, 3, 4, 2, 2], False),
            ([1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1], False),
            ([1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1, 2], True),
            ([3, 3, 3, 2, 1, 1], False),
        ]:
            output = self.isSelfCrossing(distance[:])
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
