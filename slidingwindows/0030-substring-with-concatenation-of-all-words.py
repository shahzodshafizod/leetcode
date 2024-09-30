from collections import Counter
from typing import List
import unittest

# https://leetcode.com/problems/substring-with-concatenation-of-all-words/

# python3 -m unittest slidingwindows/0030-substring-with-concatenation-of-all-words.py

class Solution(unittest.TestCase):
    # Hash Table + Sliding Window
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        n = len(s)
        words_len = len(words)
        word_len = len(words[0])
        dictionary = {word:True for word in words}
        window_len = word_len * words_len
        result = []
        for idx in range(word_len):
            count = Counter(words)
            words_found = 0
            left = idx
            for right in range(left+word_len, n+1, word_len):
                word = s[right-word_len:right]
                if word in dictionary:
                    if count[word] > 0:
                        words_found += 1
                    count[word] -= 1
                if right - left == window_len:
                    if words_found == words_len:
                        result.append(left)
                    word = s[left:left+word_len]
                    if word in dictionary:
                        count[word] += 1
                        if count[word] > 0:
                            words_found -= 1
                    left += word_len
        return result

    # # Brute-Force + Backtracking
    # def findSubstring(self, s: str, words: List[str]) -> List[int]:
    #     n = len(s)
    #     wordLen = len(words[0])
    #     count = Counter(words)
    #     def isWindow(idx: int, wordsRemained: int) -> bool:
    #         if wordsRemained == 0:
    #             return True
    #         if idx+wordLen > n:
    #             return False
    #         word = s[idx:idx+wordLen]
    #         if count.get(word, 0) == 0:
    #             return False
    #         count[word] -= 1
    #         res = isWindow(idx+wordLen, wordsRemained-1)
    #         count[word] += 1
    #         return res

    #     result = []
    #     windowLen = wordLen * len(words)
    #     for idx in range(n-windowLen+1):
    #         if isWindow(idx, len(words)):
    #             result.append(idx)
    #     return result

    def test(self) -> None:
        for s, words, expected in [
            ("barfoothefoobarman", ["foo","bar"], [0,9]),
            ("wordgoodgoodgoodbestword", ["word","good","best","word"], []),
            ("wordgoodgoodgoodbestword", ["word","good","best","good"], [8]),
            ("barfoofoobarthefoobarman", ["bar","foo","the"], [6,9,12]),
            ("lingmindraboofooowingdingbarrwingmonkeypoundcake", ["fooo","barr","wing","ding","wing"], [13]),
            ("arrarra", ["ar", "rr", "ra"], [0,1]),
            ("aaaaaaaaaaaaaa", ["aa","aa"], [0,1,2,3,4,5,6,7,8,9,10]),
            ("ooooooooooooooooooooooooooooooo", ["oo", "oo"], [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27]),
            ("aaaaa", ["aa","aa"], [0,1]),
            ("aaa", ["a", "a"], [0,1]),
            ("Ftimerunsfast", ["time","runs","fast"], [1]),
            ("barfoofoxobarfhefoobarman", ["arf","rfo","ofo","xob"], [2]),
            ("abababab", ["ab","ab","ab"], [0,2]),
            ("ababaab", ["ab","ba","ba"], [1]),
        ]:
            output = self.findSubstring(s, words)
            output.sort()
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
