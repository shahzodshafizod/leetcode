from typing import List  # , Dict, Tuple
import unittest
from collections import deque

# https://leetcode.com/problems/shortest-path-to-get-all-keys/

# python3 -m unittest graphs/0864-shortest-path-to-get-all-keys.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force DFS with memoization
    # # Time: O(m * n * 2^k * 4^(m*n)) - exponential due to exploring all paths
    # # Space: O(m * n * 2^k) for memoization
    # def shortestPathAllKeys(self, grid: List[str]) -> int:
    #     m, n = len(grid), len(grid[0])

    #     # Find start and count keys
    #     start_r, start_c = 0, 0
    #     key_count = 0

    #     for r in range(m):
    #         for c in range(n):
    #             if grid[r][c] == '@':
    #                 start_r, start_c = r, c
    #             elif 'a' <= grid[r][c] <= 'f':
    #                 key_count += 1

    #     all_keys = (1 << key_count) - 1
    #     inf = 1 << 31
    #     memo: Dict[Tuple[int, int, int], int] = {}

    #     def dfs(r: int, c: int, keys: int) -> int:
    #         if keys == all_keys:
    #             return 0

    #         if (r, c, keys) in memo:
    #             return memo[(r, c, keys)]

    #         result = inf

    #         for dr, dc in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
    #             nr, nc = r + dr, c + dc

    #             if not (0 <= nr < m and 0 <= nc < n):
    #                 continue

    #             cell = grid[nr][nc]

    #             if cell == '#':
    #                 continue

    #             new_keys = keys

    #             if 'a' <= cell <= 'f':
    #                 new_keys = keys | (1 << (ord(cell) - ord('a')))

    #             if 'A' <= cell <= 'F':
    #                 if not (keys & (1 << (ord(cell) - ord('A')))):
    #                     continue

    #             sub_result = dfs(nr, nc, new_keys)
    #             if sub_result != inf:
    #                 result = min(result, 1 + sub_result)

    #         memo[(r, c, keys)] = result
    #         return result

    #     res = dfs(start_r, start_c, 0)
    #     return res if res != inf else -1

    # Approach #2: BFS with State Compression (Bitmask)
    # Time: O(m * n * 2^k) where m,n = grid dimensions, k = number of keys
    # Space: O(m * n * 2^k) for visited set
    def shortestPathAllKeys(self, grid: List[str]) -> int:
        m, n = len(grid), len(grid[0])

        # Find start position and count total keys
        start_r, start_c = 0, 0
        key_count = 0

        for r in range(m):
            for c in range(n):
                if grid[r][c] == '@':
                    start_r, start_c = r, c
                elif 'a' <= grid[r][c] <= 'f':
                    key_count += 1

        # BFS: (row, col, keys_mask, steps)
        queue = deque([(start_r, start_c, 0, 0)])
        visited = {(start_r, start_c, 0)}
        all_keys = (1 << key_count) - 1

        directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]

        while queue:
            r, c, keys, steps = queue.popleft()

            # Check if we have all keys
            if keys == all_keys:
                return steps

            # Explore 4 directions
            for dr, dc in directions:
                nr, nc = r + dr, c + dc

                # Check bounds
                if not (0 <= nr < m and 0 <= nc < n):
                    continue

                cell = grid[nr][nc]

                # Check if wall
                if cell == '#':
                    continue

                new_keys = keys

                # Check if it's a key
                if 'a' <= cell <= 'f':
                    new_keys = keys | (1 << (ord(cell) - ord('a')))

                # Check if it's a lock
                if 'A' <= cell <= 'F':
                    # Check if we have the key
                    if not (keys & (1 << (ord(cell) - ord('A')))):
                        continue

                # Check if state already visited
                if (nr, nc, new_keys) in visited:
                    continue

                visited.add((nr, nc, new_keys))
                queue.append((nr, nc, new_keys, steps + 1))

        return -1

    def test(self):
        for grid, expected in [
            (["@.a..", "###.#", "b.A.B"], 8),
            (["@..aA", "..B#.", "....b"], 6),
            (["@Aa"], -1),
            (["@...a", ".###A", "b.BCc"], 10),
        ]:
            output = self.shortestPathAllKeys(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
