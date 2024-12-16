from typing import List
import unittest

# https://leetcode.com/problems/text-justification/

# python3 -m unittest strings/0068-text-justification.py

class Solution(unittest.TestCase):
    # Approach: Brute-Force
    # Time: O(N*maxWidth)
    # Space: O(maxWidth)
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        lines = []
        line, width = [], 0
        for idx in range(len(words)):
            if width+len(line)+len(words[idx]) > maxWidth:
                # justify
                wcount = len(line)
                spaces = maxWidth - width
                extra_spaces = spaces % max(1, wcount-1)
                spaces //= max(1, wcount-1)
                for j in range(wcount-1):
                    line[j] += ' ' * spaces
                    if extra_spaces:
                        line[j] += ' '
                        extra_spaces -= 1
                if wcount == 1:
                    line[0] += ' ' * spaces
                lines.append(''.join(line))
                line, width = [], 0 # reset
            line.append(words[idx])
            width += len(words[idx])
        last_line = ' '.join(line)
        last_line += ' ' * (maxWidth - len(last_line))
        lines.append(last_line)
        return lines

    def test(self):
        for words, maxWidth, expected in [
            (
                ["This","is","an","example","of","text","justification."],
                16,
                ["This    is    an","example  of text","justification.  "],
            ),
            (
                ["Here","is","an","example","of","text","justification."],
                14,
                ["Here   is   an","example     of","text          ","justification."],
            ),
            (
                ["What","must","be","acknowledgment","shall","be"],
                16,
                ["What   must   be","acknowledgment  ","shall be        "],
            ),
            (
                ["What","must","be","shall","be."],
                12,
                ["What must be","shall be.   "],
            ),
            (
                ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"],
                20,
                ["Science  is  what we","understand      well","enough to explain to","a  computer.  Art is","everything  else  we","do                  "],
            ),
            (
                ["My","momma","always","said,","Life","was","like","a","box","of","chocolates.","You","never","know","what","you're","gonna","get."],
                20,
                ["My    momma   always","said,  Life was like","a box of chocolates.","You  never know what","you're gonna get.   "],
            ),
            (
                ["a","b","c","d"],
                3,
                ["a b","c d"],
            ),
            (
                ["a"],
                2,
                ["a "],
            ),
            (
                ["", "", "", "","a"],
                4,
                ["    ","a   "],
            ),
        ]:
            output = self.fullJustify(words, maxWidth)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
