import unittest

# https://leetcode.com/problems/longest-duplicate-substring/

# python3 -m unittest strings/1044-longest-duplicate-substring.py


class Solution(unittest.TestCase):
    # Approach: Binary Search + Rolling Hash (Rabin-Karp Algorithm)
    # Key insight: Binary search on the length of duplicate substring.
    # For each length, use rolling hash to efficiently check if duplicate exists.
    #
    # Strategy:
    # 1. Binary search on length: if length L has duplicate, try L+1
    # 2. Rolling hash: compute hash for each substring of given length
    # 3. Store hashes in dict with indices to handle collisions
    # 4. When hash collision found, verify actual strings match
    #
    # Time Complexity: O(n log n) - binary search O(log n) × rolling hash O(n)
    # Space Complexity: O(n) for hash map
    def longestDupSubstring(self, s: str) -> str:
        def search(length: int) -> str:
            """Search for duplicate substring of given length using rolling hash"""
            base = 26         # For lowercase letters a-z
            mod = 1000000007  # Prime modulo to reduce collisions

            # Compute base^(length-1) % mod
            power = 1
            for _ in range(length - 1):
                power = (power * base) % mod

            # Compute initial hash for first window
            hash_val = 0
            for i in range(length):
                hash_val = (hash_val * base + ord(s[i]) - ord('a')) % mod

            # Map hash to list of starting indices (to handle collisions)
            seen = {hash_val: [0]}

            # Slide the window and compute rolling hash
            for i in range(length, len(s)):
                # Remove leftmost character
                hash_val = (hash_val - (ord(s[i - length]) - ord('a')) * power) % mod

                # Add rightmost character
                hash_val = (hash_val * base + ord(s[i]) - ord('a')) % mod

                # Check if this hash was seen before
                if hash_val in seen:
                    current = s[i - length + 1 : i + 1]
                    # Verify actual strings match (handle hash collisions)
                    for idx in seen[hash_val]:
                        if s[idx : idx + length] == current:
                            return current
                    # Hash collision but strings don't match
                    seen[hash_val].append(i - length + 1)
                else:
                    seen[hash_val] = [i - length + 1]

            return ""

        left, right = 1, len(s) - 1
        result = ""

        # Binary search on substring length
        while left <= right:
            mid = (left + right) // 2
            dup = search(mid)
            if dup:
                result = dup
                left = mid + 1  # Try longer length
            else:
                right = mid - 1  # Try shorter length

        return result

    def test(self):
        for s, expected in [
            ("banana", "ana"),
            ("abcd", ""),
        ]:
            output = self.longestDupSubstring(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
