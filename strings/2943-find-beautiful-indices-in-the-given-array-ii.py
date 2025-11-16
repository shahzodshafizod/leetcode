from typing import List, Tuple
import unittest

# https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/

# python3 -m unittest strings/2943-find-beautiful-indices-in-the-given-array-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-force - Simple String Matching + Linear Search
    # # Time: O(n*m + n*p + |a_indices| * |b_indices|), where n=len(s), m=len(a), p=len(b)
    # # Space: O(|a_indices| + |b_indices|)
    # def beautifulIndices(self, s: str, a: str, b: str, k: int) -> List[int]:
    #     def find_occurrences(substr: str) -> List[int]:
    #         indices: List[int] = []
    #         start = 0
    #         while True:
    #             start = s.find(substr, start)
    #             if start == -1:
    #                 break
    #             indices.append(start)
    #             start += 1
    #         return indices

    #     a_indices = find_occurrences(a)
    #     b_indices = find_occurrences(b)

    #     result: List[int] = []
    #     for i in a_indices:
    #         for j in b_indices:
    #             if abs(j - i) <= k:
    #                 result.append(i)
    #                 break

    #     return result

    # Approach #2: KMP Pattern Matching + Binary Search
    # Time: O(N+M1+M2)
    # Space: O(M1+M2)
    def beautifulIndices(self, s: str, a: str, b: str, k: int) -> List[int]:
        def kmp(substr: str) -> List[int]:
            n = len(substr)
            lps = [0] * n
            preflen, idx = 0, 1
            while idx < n:
                if substr[idx] == substr[preflen]:
                    preflen += 1
                    lps[idx] = preflen
                    idx += 1
                elif preflen > 0:
                    preflen = lps[preflen - 1]
                else:
                    lps[idx] = 0
                    idx += 1
            return lps

        na, nb = len(a), len(b)
        v1 = kmp(a + '#' + s)
        v2 = kmp(b + '#' + s)
        aindices = [i - na - na for i, v in enumerate(v1) if v >= na]
        bindices = [j - nb - nb for j, v in enumerate(v2) if v >= nb]
        res: List[int] = []
        j, nj = 0, len(bindices)
        for i in aindices:
            while j < nj and bindices[j] < i - k:
                j += 1
            if j < nj and bindices[j] <= i + k:
                res.append(i)
        return res

    def test(self):
        test_cases: List[Tuple[str, str, str, int, List[int]]] = [
            ("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15, [16, 33]),
            ("abcd", "a", "a", 4, [0]),
            ("lahhnlwx", "hhnlw", "ty", 6, []),
            ("sqgxo", "s", "xo", 3, [0]),
            ("aba", "a", "a", 1, [0, 2]),
            ("anvuyenjdp", "jd", "nu", 4, []),
        ]
        for s, a, b, k, expected in test_cases:
            output = self.beautifulIndices(s, a, b, k)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
