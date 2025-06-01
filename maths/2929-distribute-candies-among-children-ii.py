import unittest

# https://leetcode.com/problems/distribute-candies-among-children-ii/

# python3 -m unittest maths/2929-distribute-candies-among-children-ii.py

class Solution(unittest.TestCase):
    # # Approach 1: Enumeration
    # # Time: O(min(n, limit))
    # # Space: O(1)
    # def distributeCandies(self, n: int, limit: int) -> int:
    #     ways = 0
    #     for a in range(min(n, limit)+1):
    #         if n-a > 2*limit: continue
    #         ways += min(n-a, limit) - max(0, n-a-limit) + 1
    #     return ways

    # Approach 2: Inclusion-Exclusion Principle
    # Time: O(1)
    # Space: O(1)
    def distributeCandies(self, n: int, limit: int) -> int:
        def count(candies: int) -> int:
            # to avoid selecting 3 children, we add 2 more dividers
            candies += 2
            if candies < 0: return 0
            return candies * (candies-1) // 2

        # step 1: calculate number of ways to
        # distribute n candies without any limit
        ways = count(n)

        # step 2: subtract the ways where at
        # least 1 child receives more than limit
        ways -= 3 * count(n - (limit+1))

        # step 3: in the step (2), we may subtract some
        # distributions multiple times - specifically, the ones
        # where two or more children receive more than limit candies.
        ways += 3 * count(n - 2*(limit+1))

        # step 4: in the step (3), we've overcounted
        # the distributions where all three children
        # exceed the limit, so we subtract those again.
        ways -= count(n - 3*(limit+1))

        return ways

    def test(self):
        for n, limit, expected in [
            (5, 2, 3),
            (3, 3, 10),
        ]:
            output = self.distributeCandies(n, limit)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
