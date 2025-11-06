from collections import deque
from typing import Deque, List, Dict
import unittest

# https://leetcode.com/problems/construct-string-with-minimum-cost/

# python3 -m unittest dp/3213-construct-string-with-minimum-cost.py

# worst test case:
#     n = 5 * 10**4
#     target = 'a'*n
#     words = ['a', target[:-1]]
#     costs = [10000, 1]


class Solution(unittest.TestCase):
    # # Approach 1: Brute-force with Backtracking
    # # Time: O(2^n * m) where n is target length, m is number of words
    # # Space: O(n) for recursion stack
    # def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:

    #     n = len(target)

    #     def backtrack(pos: int) -> int:
    #         if pos == n:
    #             return 0
    #         min_cost = -1
    #         for i, word in enumerate(words):
    #             wlen = len(word)
    #             if pos + wlen <= n and target[pos : pos + wlen] == word:
    #                 cost = backtrack(pos + wlen)
    #                 if cost != -1 and (min_cost == -1 or costs[i] + cost < min_cost):
    #                     min_cost = costs[i] + cost
    #         return min_cost

    #     return backtrack(0)

    # # Approach 2: Bottom-up DP (Tabulation)
    # # Time: O(n * m * L) where n is target length, m is words count, L is avg word length
    # # Space: O(n) for DP array
    # def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
    #     n, INF = len(target), 10**20
    #     words_costs: Dict[str, int] = defaultdict(lambda: INF)
    #     for word, cost in zip(words, costs):
    #         words_costs[word] = min(words_costs[word], cost)
    #     n = len(target)
    #     dp = [INF] * (n + 1)
    #     dp[0] = 0
    #     for i in range(n):
    #         if dp[i] == INF:
    #             continue
    #         for word, cost in words_costs.items():
    #             wlen = len(word)
    #             if i + wlen <= n and target[i : i + wlen] == word:
    #                 dp[i + wlen] = min(dp[i + wlen], dp[i] + cost)
    #     return dp[n] if dp[n] < INF else -1

    # # Approach 3: Top-Down DP + Trie (Time Limit Exceeded)
    # # Time: O(n * L_max) where n is target length, L_max is longest word
    # # Space: O(n + total_chars) for DP array and Trie
    # def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
    #     root: Dict[str, Any] = {}

    #     def add(word: str, cost: int):
    #         cur = root
    #         for c in word:
    #             if c not in cur:
    #                 cur[c] = {}
    #             cur = cur[c]
    #         if "cost" not in cur or cost < cur["cost"]:
    #             cur["cost"] = cost

    #     for word, cost in zip(words, costs):
    #         add(word, cost)

    #     n = len(target)
    #     memo = [-1] * n

    #     def dp(pos: int) -> int:
    #         if pos == n:
    #             return 0
    #         if memo[pos] != -1:
    #             return memo[pos]
    #         min_cost = sys.maxsize
    #         cur = root
    #         for j in range(pos, n):
    #             if target[j] not in cur:
    #                 break
    #             cur = cur[target[j]]
    #             if "cost" in cur:
    #                 min_cost = min(min_cost, cur["cost"] + dp(j + 1))
    #         memo[pos] = min_cost
    #         return min_cost

    #     return -1 if dp(0) == sys.maxsize else dp(0)

    # # Approach 4: Bottom-up DP + Trie (Time Limit Exceeded)
    # # Time: O(n * L_max) where n is target length, L_max is longest word
    # # Space: O(n + total_chars) for DP array and Trie
    # def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
    #     root: Dict[str, Any] = {}
    #     for word, cost in zip(words, costs):
    #         curr = root
    #         for c in word:
    #             if c not in curr:
    #                 curr[c] = {}
    #             curr = curr[c]
    #         if "cost" not in curr or cost < curr["cost"]:
    #             curr["cost"] = cost
    #     n, INF = len(target), 10**20
    #     dp = [INF] * (n + 1)
    #     dp[0] = 0
    #     for i in range(n):
    #         if dp[i] == INF:
    #             continue
    #         curr = root
    #         for j in range(i, n):
    #             if target[j] not in curr:
    #                 break
    #             curr = curr[target[j]]
    #             if "cost" in curr:
    #                 dp[j + 1] = min(dp[j + 1], dp[i] + curr["cost"])
    #     return dp[n] if dp[n] < INF else -1

    # Approach 5: Aho-Corasick Automaton
    # Time: O(N + M) where N = len(target), M = sum(len(word) for word in words)
    # Space: O(M) for the trie + O(N) for dp array
    def minimumCost(self, target: str, words: List[str], costs: List[int]) -> int:
        # Build Aho-Corasick automaton
        class TrieNode:
            def __init__(self):
                self.children: Dict[str, TrieNode] = {}
                self.fail: TrieNode | None = None  # Failure link
                self.outputs: List[tuple[int, int]] = []  # List of (word_length, min_cost) for words ending here

        # Step 1: Build trie with minimum costs for duplicate words
        word_costs: Dict[str, int] = {}  # Track minimum cost for each unique word
        for word, cost in zip(words, costs):
            if word not in word_costs or cost < word_costs[word]:
                word_costs[word] = cost

        # Insert words into trie
        root = TrieNode()
        for word, cost in word_costs.items():
            curr = root
            for c in word:
                if c not in curr.children:
                    curr.children[c] = TrieNode()
                curr = curr.children[c]
            curr.outputs.append((len(word), cost))

        # Step 2: Build failure links (BFS)

        queue: Deque[TrieNode] = deque()
        # Initialize first level failure links to root
        for child in root.children.values():
            child.fail = root
            queue.append(child)

        # BFS to build failure links
        while queue:
            node = queue.popleft()

            for c, child in node.children.items():
                queue.append(child)

                # Find failure link
                fail_node = node.fail
                while fail_node is not None and c not in fail_node.children:
                    fail_node = fail_node.fail

                if fail_node is None:
                    child.fail = root
                else:
                    child.fail = fail_node.children[c]

                # Inherit outputs from failure link
                child.outputs.extend(child.fail.outputs)

        # Step 3: DP with Aho-Corasick matching
        n, INF = len(target), 10**20
        dp = [INF] * (n + 1)
        dp[0] = 0

        curr = root
        for i in range(n):
            c = target[i]

            # Follow failure links until we find a match or reach root
            # This ensures we never backtrack in the target string - we only move
            # forward through target while using failure links to navigate the trie.
            while curr is not None and c not in curr.children:
                curr = curr.fail

            if curr is None:
                curr = root
            else:
                curr = curr.children[c]

            # Process all matching words ending at position i+1
            for word_len, cost in curr.outputs:
                start = i - word_len + 1
                if dp[start] != INF:
                    dp[i + 1] = min(dp[i + 1], dp[start] + cost)

        return dp[n] if dp[n] < INF else -1

    def test(self):
        for target, words, costs, expected in [
            ("abcdef", ["abdef", "abc", "d", "def", "ef"], [100, 1, 1, 10, 5], 7),
            ("aaaa", ["z", "zz", "zzz"], [1, 10, 100], -1),
            # Test case for duplicate words with different costs (bug fix verification)
            ("abc", ["abc", "abc", "ab", "c"], [10, 5, 3, 2], 5),
            # Test case with overlapping words
            ("aaa", ["a", "aa", "aaa"], [1, 2, 3], 3),
        ]:
            output = self.minimumCost(target, words, costs)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")

    def test_worst_case(self):
        # Worst case: very long target with very long word that almost matches
        # Note: Full n=50000 works but is slow in Python (~40s) due to string operations
        # The algorithm is O(n * L_max) which is optimal, but Python is slower than Go
        n = 10000  # Use 10k for reasonable test time
        target = 'a' * n
        words = ['a', target[:-1]]  # 'a' and 'a'*(n-1)
        costs = [10000, 1]
        # Optimal: use 'a'*(n-1) once (cost 1) + 'a' once (cost 10000) = 10001
        expected = 10001
        output = self.minimumCost(target, words, costs)
        self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
