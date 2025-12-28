from typing import List, Tuple
import unittest

# https://leetcode.com/problems/maximum-profit-from-trading-stocks-with-discounts/

# python3 -m unittest dp/3562-maximum-profit-from-trading-stocks-with-discounts.py


class Solution(unittest.TestCase):
    # Approach: Tree DP with Bounded Knapsack (Optimized)
    # Intuition: Use tree DP where for each node we track:
    # - dp0[cost] = max profit in subtree using exact cost when parent didn't buy
    # - dp1[cost] = max profit in subtree using exact cost when parent bought
    # We use bounded arrays sized by maximum possible cost in each subtree.
    # Use -1 to mark invalid states.
    # Time Complexity: O(n * budget^2)
    # Space Complexity: O(n * budget)
    def maximumProfit(
        self, n: int, present: List[int], future: List[int], hierarchy: List[List[int]], budget: int
    ) -> int:
        # Build adjacency list
        children: List[List[int]] = [[] for _ in range(n + 1)]
        for u, v in hierarchy:
            children[u].append(v)

        def dfs(u: int) -> Tuple[List[int], List[int], int]:
            cost = present[u - 1]
            d_cost = cost // 2
            profit = future[u - 1] - cost
            d_profit = future[u - 1] - d_cost

            # subProfit0[w] = max profit from children with exact cost w, children see unbought parent
            # subProfit1[w] = max profit from children with exact cost w, children see bought parent
            sub_profit0 = [0]
            sub_profit1 = [0]

            u_size = cost
            for v in children[u]:
                child_dp0, child_dp1, child_size = dfs(v)
                u_size += child_size

                # Merge with child using bounded knapsack
                new_len = min(budget + 1, len(sub_profit0) + len(child_dp0) - 1)
                new_sub0 = [-1] * new_len

                for i in range(len(sub_profit0)):
                    if i >= new_len or sub_profit0[i] == -1:
                        continue
                    for j in range(len(child_dp0)):
                        if i + j >= new_len or child_dp0[j] == -1:
                            continue
                        new_sub0[i + j] = max(new_sub0[i + j], sub_profit0[i] + child_dp0[j])

                sub_profit0 = new_sub0

                new_len = min(budget + 1, len(sub_profit1) + len(child_dp1) - 1)
                new_sub1 = [-1] * new_len

                for i in range(len(sub_profit1)):
                    if i >= new_len or sub_profit1[i] == -1:
                        continue
                    for j in range(len(child_dp1)):
                        if i + j >= new_len or child_dp1[j] == -1:
                            continue
                        new_sub1[i + j] = max(new_sub1[i + j], sub_profit1[i] + child_dp1[j])

                sub_profit1 = new_sub1

            # Build dp0: parent didn't buy
            max_len0 = min(budget + 1, max(len(sub_profit0), len(sub_profit1) + cost))
            dp0 = [-1] * max_len0

            # Option 1: don't buy u
            for i in range(min(len(sub_profit0), max_len0)):
                dp0[i] = max(dp0[i], sub_profit0[i])

            # Option 2: buy u at full cost
            if cost < max_len0:
                for i in range(len(sub_profit1)):
                    if i + cost < max_len0 and sub_profit1[i] != -1:
                        dp0[i + cost] = max(dp0[i + cost], sub_profit1[i] + profit)

            # Build dp1: parent bought (discount available)
            max_len1 = min(budget + 1, max(len(sub_profit0), len(sub_profit1) + d_cost))
            dp1 = [-1] * max_len1

            # Option 1: don't buy u
            for i in range(min(len(sub_profit0), max_len1)):
                dp1[i] = max(dp1[i], sub_profit0[i])

            # Option 2: buy u at discounted cost
            if d_cost < max_len1:
                for i in range(len(sub_profit1)):
                    if i + d_cost < max_len1 and sub_profit1[i] != -1:
                        dp1[i + d_cost] = max(dp1[i + d_cost], sub_profit1[i] + d_profit)

            return dp0, dp1, u_size

        dp0, _, _ = dfs(1)
        return max(dp0[i] for i in range(min(budget + 1, len(dp0))) if dp0[i] != -1)

    def test(self):
        test_cases: List[Tuple[int, List[int], List[int], List[List[int]], int, int]] = [
            (2, [1, 2], [4, 3], [[1, 2]], 3, 5),
            (2, [3, 4], [5, 8], [[1, 2]], 4, 4),
            (3, [4, 6, 8], [7, 9, 11], [[1, 2], [1, 3]], 10, 10),
            (3, [5, 2, 3], [8, 5, 6], [[1, 2], [2, 3]], 7, 12),
            (1, [45], [38], [], 115, 0),
            (2, [6,11], [5,48], [[1,2]], 142, 42),
            (3, [6,4,23], [50,48,17], [[1,3],[1,2]], 28, 96),
            (3, [42,27,32], [46,8,17], [[1,2],[2,3]], 93, 4),
        ]
        for n, present, future, hierarchy, budget, expected in test_cases:
            output = self.maximumProfit(n, present, future, hierarchy, budget)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
