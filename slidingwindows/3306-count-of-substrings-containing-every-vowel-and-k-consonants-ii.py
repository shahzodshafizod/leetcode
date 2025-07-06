import unittest

# https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/

# python3 -m unittest slidingwindows/3306-count-of-substrings-containing-every-vowel-and-k-consonants-ii.py


class Solution(unittest.TestCase):
    def countOfSubstrings(self, word: str, k: int) -> int:
        n = len(word)

        def at_least(k: int) -> int:
            vowels = set(['a', 'e', 'i', 'o', 'u'])
            result = 0
            count = [0] * 26
            vcount, ccount = 0, 0
            left = 0
            for right in range(n):
                code = ord(word[right]) - ord('a')
                count[code] += 1
                if word[right] not in vowels:
                    ccount += 1
                elif count[code] == 1:
                    vcount += 1
                while vcount == 5 and ccount >= k:
                    result += n - right
                    code = ord(word[left]) - ord('a')
                    count[code] -= 1
                    if word[left] not in vowels:
                        ccount -= 1
                    elif count[code] == 0:
                        vcount -= 1
                    left += 1
            return result

        return at_least(k) - at_least(k + 1)

    def test(self):
        for word, k, expected in [
            ("aeioqq", 1, 0),
            ("aeiou", 0, 1),
            ("ieaouqqieaouqq", 1, 3),
            ("iqeaouqi", 2, 3),
            ("aeouih", 0, 1),
            ("aadieuoh", 1, 2),
            ("euaoei", 1, 0),
            ("aoaiuefi", 1, 4),
        ]:
            output = self.countOfSubstrings(word, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
