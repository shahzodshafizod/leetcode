from typing import List
from collections import Counter
import unittest

# https://leetcode.com/problems/shortest-completing-word/

# python3 -m unittest hashes/0748-shortest-completing-word.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # For each word, manually count character frequencies and compare
    # # Time: O(n * m) where n is number of words, m is average word length
    # # Space: O(1) - fixed size array for 26 letters
    # def shortestCompletingWord(self, licensePlate: str, words: List[str]) -> str:
    #     # Count letters in license plate (ignore case, numbers, spaces)
    #     def count_letters(s: str) -> List[int]:
    #         counts = [0] * 26
    #         for char in s.lower():
    #             if char.isalpha():
    #                 counts[ord(char) - ord('a')] += 1
    #         return counts

    #     # Check if word contains all letters from license plate
    #     def is_completing(word: str, target_counts: List[int]) -> bool:
    #         word_counts = count_letters(word)
    #         for i in range(26):
    #             if word_counts[i] < target_counts[i]:
    #                 return False
    #         return True

    #     target_counts = count_letters(licensePlate)
    #     result: str = ""
    #     min_length = float('inf')

    #     for word in words:
    #         if is_completing(word, target_counts):
    #             if len(word) < min_length:
    #                 min_length = len(word)
    #                 result = word

    #     return result

    # Approach 2: Optimized using Counter (Hash Table)
    # Use Counter from collections for cleaner frequency counting
    # Time: O(n * m) where n is number of words, m is average word length
    # Space: O(1) - Counter stores at most 26 letters
    def shortestCompletingWord(self, licensePlate: str, words: List[str]) -> str:
        # Extract and count only letters from license plate
        plate_letters = ''.join(c.lower() for c in licensePlate if c.isalpha())
        plate_counter = Counter(plate_letters)

        result: str = ""
        min_length = float('inf')

        for word in words:
            # Check if this word contains all required letters
            word_counter = Counter(word)

            # Check if word has all letters with required frequencies
            is_valid = all(word_counter[char] >= count for char, count in plate_counter.items())

            if is_valid and len(word) < min_length:
                min_length = len(word)
                result = word

        return result

    def test(self):
        for licensePlate, words, expected in [
            ("1s3 PSt", ["step", "steps", "stripe", "stepple"], "steps"),
            ("1s3 456", ["looks", "pest", "stew", "show"], "pest"),
            ("Ah71752", ["suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"], "husband"),
            ("OgEu755", ["enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"], "enough"),
        ]:
            output = self.shortestCompletingWord(licensePlate, words)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
