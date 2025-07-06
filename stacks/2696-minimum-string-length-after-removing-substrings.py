import unittest

# https://leetcode.com/problems/minimum-string-length-after-removing-substrings/

# python3 -m unittest stacks/2696-minimum-string-length-after-removing-substrings.py


class Solution(unittest.TestCase):
    # Approach: Stack
    # Time: O(N)
    # Space: O(N)
    def minLength(self, s: str) -> int:
        stack = []
        for c in s:
            if stack and c in ('B', 'D') and ord(stack[-1]) + 1 == ord(c):
                stack.pop()
            else:
                stack.append(c)
        return len(stack)

    def test(self):
        for s, expected in [
            ("ABFCACDB", 2),
            ("ACBBD", 5),
            ("CCADDADDDBBCDDBABACADABAABADCABDCCBDACBDBAADDABCABBCABBDDAABCBCBBCCCDBCDDDADAACBCACDDBBA", 62),
            ("DCDCBCBDACBBABDABABCCCBABCCCCCCCBDDBCDACDADABADDCDBBC", 35),
            ("BBBDCADCDACACDBBCACDACDABCBCDDADCDCACCDDBCACCDDDCCBCDBDCBDDCBCBBCCBBBAACBBB", 47),
            ("BCDDBCCCCBACCADDCBDADDCCABCCCBACAADDADCDAACABDDDDABBACBABCABCCDCABBA", 48),
            ("RZAAAACCCCCAABBDDDDDBBBBYAAAAAAACAACAACACCAAAACCACCDDBDDBBBBDDBDBBDBBDBBBBBBBMBSCACCACDBDDBDACCDDBCD", 6),
            ("DCAACCCCCCACCCAACAABBDBBDDDBDDDDDDBBDPCAACAAACAAAAAACDBBBBBBDBBBDBBDACCCAAABBBDDDAACACDBDBBCAACDBBDP", 4),
            ("AXFCCCCCACCACCCAACAACACACAACCAAABBBDDBBDBDBDBBDBBDDDBDDBDDDDDCCACDBDDCABD", 3),
            ("ACACACACACACACACACACACACACACACACACACACACACACACACACDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDBDB", 0),
        ]:
            output = self.minLength(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
