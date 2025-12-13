import unittest

# https://leetcode.com/problems/reverse-vowels-of-a-string/

# python3 -m unittest strings/0345-reverse-vowels-of-a-string.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Find all vowels, reverse them, rebuild string
    # # Extract vowel positions and values, reverse values, rebuild.
    # # Time: O(n)
    # # Space: O(n)
    # def reverseVowels(self, s: str) -> str:
    #     vowels = set('aeiouAEIOU')
    #     vowel_chars = [c for c in s if c in vowels]
    #     vowel_chars.reverse()

    #     result: List[str] = []
    #     vowel_idx = 0
    #     for c in s:
    #         if c in vowels:
    #             result.append(vowel_chars[vowel_idx])
    #             vowel_idx += 1
    #         else:
    #             result.append(c)

    #     return ''.join(result)

    # Approach #2: Two Pointers
    # Use two pointers from both ends, swap when both point to vowels.
    # Time: O(n) - single pass
    # Space: O(n) - convert to list for mutation
    def reverseVowels(self, s: str) -> str:
        vowels = set('aeiouAEIOU')
        chars = list(s)
        left, right = 0, len(chars) - 1

        while left < right:
            if chars[left] not in vowels:
                left += 1
            elif chars[right] not in vowels:
                right -= 1
            else:
                chars[left], chars[right] = chars[right], chars[left]
                left += 1
                right -= 1

        return ''.join(chars)

    def test(self):
        for s, expected in [
            ("IceCreAm", "AceCreIm"),
            ("leetcode", "leotcede"),
            ("hello", "holle"),
            ("aA", "Aa"),
        ]:
            output = self.reverseVowels(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
