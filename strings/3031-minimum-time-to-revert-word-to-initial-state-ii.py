from typing import List
import unittest

# https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/

# python3 -m unittest strings/3031-minimum-time-to-revert-word-to-initial-state-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # After each operation, remove first k chars and check if remaining can form original
    # # Time: O(n^2 / k) - too slow due to repeated string prefix comparisons
    # # This approach uses the built-in startswith which performs O(n) comparison at each step,
    # # leading to quadratic behavior for the worst case
    # # Space: O(n) for string slicing
    # def minimumTimeToInitialState(self, word: str, k: int) -> int:
    #     n = len(word)
    #     steps = 1
    #     pos = k

    #     while pos < n:
    #         # Check if word[pos:] is a prefix of word
    #         if word.startswith(word[pos:]):
    #             return steps
    #         pos += k
    #         steps += 1

    #     return steps

    # Approach #2: Optimized - Use Z-algorithm or KMP for efficient prefix matching
    # Z-algorithm: For each position i, z[i] is length of longest substring starting at i
    # that matches a prefix of the string
    # Time: O(n) for Z-algorithm construction
    # Space: O(n) for Z array
    def minimumTimeToInitialState(self, word: str, k: int) -> int:
        n = len(word)

        # Build Z-array
        def z_algorithm(s: str) -> List[int]:
            n = len(s)
            z = [0] * n
            z[0] = n
            left, right = 0, 0

            for i in range(1, n):
                if i > right:
                    left = right = i
                    while right < n and s[right - left] == s[right]:
                        right += 1
                    z[i] = right - left
                    right -= 1
                else:
                    k_pos = i - left
                    if z[k_pos] < right - i + 1:
                        z[i] = z[k_pos]
                    else:
                        left = i
                        while right < n and s[right - left] == s[right]:
                            right += 1
                        z[i] = right - left
                        right -= 1

            return z

        z = z_algorithm(word)

        # Check positions at multiples of k
        steps = 1
        pos = k
        while pos < n:
            # If remaining length matches prefix
            if z[pos] >= n - pos:
                return steps
            pos += k
            steps += 1

        return steps

    def test(self):
        for word, k, expected in [
            ("abacaba", 3, 2),
            ("abacaba", 4, 1),
            ("aabaa", 2, 2),
            ("abc", 1, 3),
        ]:
            output = self.minimumTimeToInitialState(word, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
