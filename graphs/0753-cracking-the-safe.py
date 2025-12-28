from collections import defaultdict
from typing import Set, List, Dict
import unittest

# https://leetcode.com/problems/cracking-the-safe/

# python3 -m unittest graphs/0753-cracking-the-safe.py

"""
Problem: Cracking the Safe

This is an Eulerian Circuit problem in disguise. The key insight is:
- Each n-digit password can be thought of as a node in a graph
- We can transition from one password to another by removing the first digit
  and adding a new digit at the end
- We need to visit all possible n-digit passwords exactly once (all edges)
- This forms an Eulerian Circuit/Path problem

For n=2, k=2:
- We have 2^2 = 4 possible passwords: 00, 01, 10, 11
- We need a string that contains all these as substrings

The graph structure:
- Nodes represent (n-1)-digit sequences
- Edges represent n-digit sequences
- Each node has exactly k outgoing edges (adding digits 0 to k-1)
- Total edges = k^n (all possible n-digit passwords)
- We need to find an Eulerian path that visits every edge exactly once
"""


class Solution(unittest.TestCase):
    # # Approach 1: Backtracking (Brute Force)
    # # Time: O(k^n * k^n) - we try to visit k^n nodes, each attempt can take k^n operations
    # # Space: O(k^n) - for visited set and recursion stack
    # # This approach tries all possible sequences and backtracks when stuck
    # def crackSafe(self, n: int, k: int) -> str:
    #     # Edge case: n=1, just return all digits
    #     if n == 1:
    #         return "".join(str(i) for i in range(k))

    #     # Start with n-1 zeros (representing the initial state)
    #     start = "0" * (n - 1)
    #     visited: Set[str] = set()
    #     result = [start]

    #     def dfs(node: str) -> bool:
    #         # Total number of passwords to visit
    #         if len(visited) == k ** n:
    #             return True

    #         # Try appending each digit from 0 to k-1
    #         for digit in range(k):
    #             # Form the new password by taking last n-1 chars + new digit
    #             password = node + str(digit)

    #             if password not in visited:
    #                 visited.add(password)
    #                 result.append(str(digit))

    #                 # The new node is the last n-1 characters of the password
    #                 new_node = password[1:]

    #                 if dfs(new_node):
    #                     return True

    #                 # Backtrack
    #                 visited.remove(password)
    #                 result.pop()

    #         return False

    #     dfs(start)
    #     return "".join(result)

    # Approach 2: Hierholzer's Algorithm (Optimal - Eulerian Circuit)
    # Time: O(k^n) - we visit each of the k^n edges exactly once
    # Space: O(k^n) - for the graph and result string
    #
    # Algorithm:
    # 1. Build a graph where nodes are (n-1)-digit strings
    # 2. Each node has k outgoing edges (one for each digit 0 to k-1)
    # 3. Use Hierholzer's algorithm to find Eulerian path
    # 4. Start from "000...0" (n-1 zeros)
    #
    # Why this works:
    # - Each node has exactly k incoming and k outgoing edges (Eulerian circuit exists)
    # - We want to traverse every edge (n-digit password) exactly once
    # - Hierholzer's algorithm efficiently finds such a path
    def crackSafe(self, n: int, k: int) -> str:
        # Edge case: if n=1, just return all digits
        if n == 1:
            return "".join(str(i) for i in range(k))

        # Build adjacency list as a set of available edges from each node
        graph: Dict[str, List[int]] = defaultdict(list)

        # Total number of (n-1)-digit nodes
        # Each node has k outgoing edges
        total_nodes = k ** (n - 1)

        # Build the graph
        for i in range(total_nodes):
            # Convert number to (n-1)-digit string
            node = ""
            num = i
            for _ in range(n - 1):
                node = str(num % k) + node
                num //= k

            # Add edges for all possible next digits
            for digit in range(k):
                graph[node].append(digit)

        # Start from all zeros
        start = "0" * (n - 1)
        stack = [start]
        path: List[str] = []

        # Hierholzer's algorithm
        while stack:
            node = stack[-1]
            if graph[node]:
                # Take an edge
                digit = graph[node].pop()
                # Next node is current node shifted left + new digit
                next_node = node[1:] + str(digit)
                stack.append(next_node)
            else:
                # No more edges from this node, add to path
                path.append(stack.pop())

        # Reverse the path (Hierholzer's algorithm builds it backwards)
        path.reverse()

        # Build result: start with first node, then append last digit of each subsequent node
        result = path[0]
        for i in range(1, len(path)):
            # Append the digit that was added to transition to this node
            result += path[i][-1]

        return result

    def test(self):
        for n, k in [
            (1, 2),
            (2, 2),
            (2, 3),
            (3, 2),
        ]:
            # Test both approaches
            output = self.crackSafe(n, k)
            self.assertTrue(self._verify(output, n, k), f"Optimized failed: n={n}, k={k}, output={output}")

    def _verify(self, s: str, n: int, k: int) -> bool:
        """Verify that the solution contains all possible n-digit passwords"""
        # Check minimum length
        expected_len = k ** n + n - 1
        if len(s) != expected_len:
            return False

        # Generate all possible passwords
        passwords: Set[str] = set()
        for i in range(k ** n):
            password = ""
            num = i
            for _ in range(n):
                password = str(num % k) + password
                num //= k
            passwords.add(password)

        # Check all passwords appear as substrings
        found: Set[str] = set()
        for i in range(len(s) - n + 1):
            found.add(s[i:i + n])

        return found == passwords
