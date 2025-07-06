from typing import List
import unittest

# https://leetcode.com/problems/remove-all-occurrences-of-a-substring/

# python3 -m unittest stacks/1910-remove-all-occurrences-of-a-substring.py


class Solution(unittest.TestCase):
    # Approach: KMP (Knuth-Morris-Pratt) Algorithm
    # Time: O(n+m), n=len(s), m=len(part)
    # Space: O(n+m)
    def removeOccurrences(self, s: str, part: str) -> str:
        pn = len(part)

        def get_lps(part: str) -> List[int]:
            lps = [0] * pn
            preflen, idx = 0, 1
            while idx < pn:
                if part[idx] == part[preflen]:
                    preflen += 1
                    lps[idx] = preflen
                    idx += 1
                elif preflen > 0:
                    preflen = lps[preflen - 1]
                else:
                    lps[idx] = 0
                    idx += 1
            return lps

        lps = get_lps(part)
        pid, sid, sn = 0, 0, len(s)
        diff = 0
        stack, nextid = [None] * sn, [0] * sn
        while sid < sn:
            stack[sid - diff] = s[sid]
            if s[sid] == part[pid]:
                pid += 1
                nextid[sid - diff] = pid
                if pid == pn:
                    diff += pn
                    pid = nextid[sid - diff] if sid >= diff else 0
            elif pid > 0:
                pid = lps[pid - 1]
                sid -= 1
            else:
                nextid[sid - diff] = 0
            sid += 1
        return "".join(stack[: sn - diff])

    def test(self):
        for s, part, expected in [
            ("a", "aa", "a"),
            ("eemckxmckx", "emckx", ""),
            ("ccctltctlltlb", "ctl", "b"),
            ("yjyjqnaxlbqnaxlbfss", "yjqnaxlb", "fss"),
            ("gjzgbpggjzgbpgsvpwdk", "gjzgbpg", "svpwdk"),
            ("hhvhvaahvahvhvaavhvaasshvahvaln", "hva", "ssln"),
            ("daabcbaabcbc", "abc", "dab"),
            ("axxxxyyyyb", "xy", "ab"),
            ("wwwwwwwwwwwwwwwwwwwwwvwwwwswxwwwwsdwxweeohapwwzwuwajrnogb", "w", "vsxsdxeeohapzuajrnogb"),
            # ("kpygkivtlqoockpygkivtlqoocssnextkqzjpycbylkaondsskpygkpygkivtlqoocssnextkqzjpkpygkivtlqoocssnextkqzjpycbylkaondsycbylkaondskivtlqoocssnextkqzjpycbylkaondssnextkqzjpycbylkaondshijzgaovndkjiiuwjtcpdpbkrfsi", "kpygkivtlqoocssnextkqzjpycbylkaonds", "hijzgaovndkjiiuwjtcpdpbkrfsi"), # pylint: disable=line-too-long
        ]:
            output = self.removeOccurrences(s, part)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
