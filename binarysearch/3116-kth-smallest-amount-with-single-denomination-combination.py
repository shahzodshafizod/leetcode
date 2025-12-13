from typing import List
import unittest

# https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/

# python3 -m unittest binarysearch/3116-kth-smallest-amount-with-single-denomination-combination.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force with Heap (Multi-way Merge)
    # # Generate multiples of each coin and merge using min-heap
    # # Time: O(k * log n) - n pointers for n coins
    # # Space: O(n) - heap size
    # def findKthSmallest(self, coins: List[int], k: int) -> int:
    #     import heapq

    #     # Min-heap: (value, coin_index, multiplier)
    #     heap = [(coin, i, 1) for i, coin in enumerate(coins)]
    #     heapq.heapify(heap)

    #     prev = 0
    #     count = 0

    #     while heap and count < k:
    #         val, coin_idx, mult = heapq.heappop(heap)

    #         # Skip duplicates
    #         if val == prev:
    #             heapq.heappush(heap, (coins[coin_idx] * (mult + 1), coin_idx, mult + 1))
    #             continue

    #         prev = val
    #         count += 1

    #         if count == k:
    #             return val

    #         # Add next multiple of this coin
    #         heapq.heappush(heap, (coins[coin_idx] * (mult + 1), coin_idx, mult + 1))

    #     return prev

    # Approach #2: Binary Search + Inclusion-Exclusion Principle
    # Binary search on the answer. For a given value x, count how many amounts <= x
    # can be formed using the coins (using inclusion-exclusion with LCM).
    # Time: O(log(k*max_coin) * 2^n) where n = len(coins)
    # Space: O(n) for recursion
    def findKthSmallest(self, coins: List[int], k: int) -> int:
        from math import gcd

        def lcm(a: int, b: int) -> int:
            return a * b // gcd(a, b)

        # Remove coins that are multiples of other coins (optimization)
        coins_set = set(coins)
        filtered_coins: List[int] = []
        for coin in sorted(coins_set):
            is_multiple = False
            for other in filtered_coins:
                if coin % other == 0:
                    is_multiple = True
                    break
            if not is_multiple:
                filtered_coins.append(coin)

        coins = filtered_coins
        n = len(coins)

        def count_amounts(x: int) -> int:
            # Count amounts <= x using inclusion-exclusion
            total = 0

            # Iterate through all subsets
            for mask in range(1, 1 << n):
                subset_lcm = 1
                bits = 0

                for i in range(n):
                    if mask & (1 << i):
                        subset_lcm = lcm(subset_lcm, coins[i])
                        bits += 1
                        # Optimization: if LCM becomes too large, skip
                        if subset_lcm > x:
                            break

                if subset_lcm <= x:
                    # Inclusion-exclusion: add if odd number of coins, subtract if even
                    if bits % 2 == 1:
                        total += x // subset_lcm
                    else:
                        total -= x // subset_lcm

            return total

        # Binary search for the kth smallest amount
        left, right = 1, k * max(coins)

        while left < right:
            mid = (left + right) // 2
            if count_amounts(mid) < k:
                left = mid + 1
            else:
                right = mid

        return left

    def test(self):
        for coins, k, expected in [
            ([3, 6, 9], 3, 9),  # Amounts: 3, 6, 9, 12, ... -> 3rd is 9
            ([5, 2], 7, 12),  # Amounts: 2, 4, 5, 6, 8, 10, 12, ... -> 7th is 12
            ([1], 5, 5),  # Amounts: 1, 2, 3, 4, 5, ... -> 5th is 5
        ]:
            output = self.findKthSmallest(coins, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
