from typing import List
import unittest

# https://leetcode.com/problems/palindrome-pairs/

# python3 -m unittest tries/0336-palindrome-pairs.py


class Solution(unittest.TestCase):
    # Approach: HashMap with Reversed Words and Palindrome Checking
    # Key insight: For words[i] + words[j] to be a palindrome, we need:
    # - Case 1: reverse(words[i]) == words[j] (equal length palindromes)
    # - Case 2: words[i] = prefix + suffix, where suffix is palindrome and reverse(prefix) exists
    # - Case 3: words[j] = prefix + suffix, where prefix is palindrome and reverse(suffix) exists
    #
    # Algorithm:
    # 1. Store all words in a hashmap with their reversed form as key
    # 2. For each word, split it into all possible (prefix, suffix) combinations
    # 3. Check:
    #    - If suffix is palindrome and reverse(prefix) exists in map → valid pair
    #    - If prefix is palindrome and reverse(suffix) exists in map → valid pair
    # 4. Handle edge cases with empty strings
    #
    # Example: words = ["abcd", "dcba", "lls", "s", "sssll"]
    # - "abcd" + "dcba" = "abcddcba" (palindrome)
    # - "lls": split into ("", "lls"), ("l", "ls"), ("ll", "s"), ("lls", "")
    #   - ("ll", "s"): "s" is palindrome, reverse("ll") = "ll" not found
    #   - Check reverse: ("s", "ll"): "s" is palindrome, reverse("ll") = "ll" not found
    #   - But checking "s" alone: reverse("s") = "s", prefix "ll" is palindrome → "lls" + "s" = "llss" (not palindrome)
    #   - Actually for "lls" + "s": need "s" + something to make palindrome with "lls"
    #
    # Time: O(n * k^2) where n is number of words, k is average word length
    #       For each word (n), we check all splits (k) and verify palindrome (k)
    # Space: O(n * k) for the hashmap storing all words
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        def is_palindrome(s: str) -> bool:
            return s == s[::-1]

        word_map = {word: i for i, word in enumerate(words)}
        result: List[List[int]] = []

        for i, word in enumerate(words):
            # Check all possible splits of current word
            for j in range(len(word) + 1):
                prefix = word[:j]
                suffix = word[j:]

                # Case 1: suffix is palindrome, check if reverse(prefix) exists
                if is_palindrome(suffix):
                    reversed_prefix = prefix[::-1]
                    if reversed_prefix in word_map and word_map[reversed_prefix] != i:
                        result.append([i, word_map[reversed_prefix]])

                # Case 2: prefix is palindrome, check if reverse(suffix) exists
                # Avoid duplicate when j == len(word) (empty suffix creates same pair as Case 1 with j==0)
                if j > 0 and is_palindrome(prefix):
                    reversed_suffix = suffix[::-1]
                    if reversed_suffix in word_map and word_map[reversed_suffix] != i:
                        result.append([word_map[reversed_suffix], i])

        return result

    def test(self):
        for words, expected in [
            (["abcd", "dcba", "lls", "s", "sssll"], [[0, 1], [1, 0], [3, 2], [2, 4]]),
            (["bat", "tab", "cat"], [[0, 1], [1, 0]]),
            (["a", ""], [[0, 1], [1, 0]]),
            (["a", "abc", "aba", ""], [[0, 3], [2, 3], [3, 0], [3, 2]]),
        ]:
            output = self.palindromePairs(words[:])
            # Sort for comparison since order may vary
            output.sort()
            expected.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
