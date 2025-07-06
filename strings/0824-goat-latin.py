import unittest

# https://leetcode.com/problems/goat-latin/

# python3 -m unittest strings/0824-goat-latin.py


class Solution(unittest.TestCase):
    def toGoatLatin(self, sentence: str) -> str:
        vowels = set(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'])

        def proc(word: str, cnt: int) -> str:
            if word[0] not in vowels:
                word = word[1:] + word[0]
            return word + "ma" + "a" * cnt

        return " ".join(proc(w, i) for i, w in enumerate(sentence.split(), 1))

    def test(self):
        for sentence, expected in [
            ("I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"),
            (
                "The quick brown fox jumped over the lazy dog",
                "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
            ),
        ]:
            output = self.toGoatLatin(sentence)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
