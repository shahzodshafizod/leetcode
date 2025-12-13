from typing import List
import unittest

# https://leetcode.com/problems/count-connected-components-in-lcm-graph/

# python3 -m unittest unionfinds/3378-count-connected-components-in-lcm-graph.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # For each pair of numbers, compute LCM and connect if LCM <= threshold
    # # Time: O(n^2 * log(max(nums))) - too slow for large n due to checking all pairs
    # # This approach requires checking every possible pair of numbers, which becomes
    # # prohibitively expensive as the array size grows
    # # Space: O(n) for union find structure
    # def countComponents(self, nums: List[int], threshold: int) -> int:
    #     def gcd(a: int, b: int) -> int:
    #         while b:
    #             a, b = b, a % b
    #         return a

    #     def lcm(a: int, b: int) -> int:
    #         return a * b // gcd(a, b)

    #     n = len(nums)
    #     parent = list(range(n))

    #     def find(x: int) -> int:
    #         if parent[x] != x:
    #             parent[x] = find(parent[x])
    #         return parent[x]

    #     def union(x: int, y: int) -> bool:
    #         px = find(x)
    #         py = find(y)
    #         if px == py:
    #             return False
    #         parent[py] = px
    #         return True

    #     # Check all pairs
    #     for i in range(n):
    #         for j in range(i + 1, n):
    #             if lcm(nums[i], nums[j]) <= threshold:
    #                 union(i, j)

    #     # Count distinct components
    #     components = len(set(find(i) for i in range(n)))
    #     return components

    # Approach #2: Union-Find with Multiples Connection
    # Key insight: For each value that exists in nums and is <= threshold,
    # connect it with all its multiples (2*val, 3*val, ...) that are also <= threshold.
    # This works because LCM(a, k*a) = k*a, so if k*a <= threshold, they connect.
    # Time: O(threshold * log(threshold) + n) - harmonic series sum
    # Space: O(threshold) for DSU structure
    def countComponents(self, nums: List[int], threshold: int) -> int:
        # DSU operates on values up to threshold
        parent = list(range(threshold + 1))

        def find(x: int) -> int:
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x: int, y: int):
            px = find(x)
            py = find(y)
            if px != py:
                parent[py] = px

        # For each value in nums that's <= threshold, connect it with all multiples
        # Time complexity: O(threshold/1 + threshold/2 + ... + threshold/k) = O(threshold * log(threshold))
        for num in nums:
            if num <= threshold:
                # Connect num with all its multiples: 2*num, 3*num, etc.
                for multiple in range(2 * num, threshold + 1, num):
                    union(num, multiple)

        # Count distinct components
        components: set[int] = set()
        large_count = 0
        for num in nums:
            if num > threshold:
                # Each number > threshold is its own component
                large_count += 1
            else:
                components.add(find(num))

        return len(components) + large_count

    def test(self):
        for nums, threshold, expected in [
            ([2, 4, 8, 3, 9], 5, 4),
            ([2, 4, 8, 3, 9, 12], 10, 2),
            ([1], 1, 1),
            ([5, 10, 15], 20, 1),
        ]:
            output = self.countComponents(nums, threshold)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
