from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/naming-a-company/

# python3 -m unittest hashes/2306-naming-a-company.py


class Solution(unittest.TestCase):
    # # Approach: Tries
    # # Time: O(NxW), N=len(ideas), W=max(len(ideas[i]))
    # # Space: O(N)
    # def distinctNames(self, ideas: List[str]) -> int:
    #     class Node:
    #         def __init__(self):
    #             self.children = {}
    #             self.start = False
    #     orda = ord('a')
    #     counts = [0] * 26
    #     same = [[0] * 26 for _ in range(26)]
    #     root = Node()
    #     for idea in ideas:
    #         curr = root
    #         prev = curr
    #         for c in range(len(idea)-1, -1, -1):
    #             key = ord(idea[c]) - orda
    #             if key not in curr.children:
    #                 curr.children[key] = Node()
    #             prev = curr
    #             curr = curr.children[key]
    #         counts[key] += 1
    #         for c, next in prev.children.items():
    #             if next.start:
    #                 same[key][c] += 1
    #                 same[c][key] += 1
    #         curr.start = True
    #     count = 0
    #     for a in range(26):
    #         for b in range(a+1, 26):
    #             counta = counts[a] - same[a][b]
    #             countb = counts[b] - same[b][a]
    #             count += counta * countb # * 2
    #     return count * 2

    # Approach: Hash Set
    # Time: O(N)
    # Space: O(N)
    def distinctNames(self, ideas: List[str]) -> int:
        count = defaultdict(set)  # S: O(26)
        for a in ideas:  # T: O(N)
            count[a[0]].add(hash(a[1:]))
        res = 0
        for a, seta in count.items():  # T: O(26)
            for b, setb in count.items():  # T: O(26)
                if b <= a:
                    continue
                same = len(seta & setb)
                res += (len(seta) - same) * (len(setb) - same)
        return res * 2

    def test(self):
        for ideas, expected in [
            (["coffee", "donuts", "time", "toffee"], 6),
            (["lack", "back"], 0),
            (["a", "b", "baby", "aby", "banana", "apple"], 8),
            # (["j","t","fmucd","liantvj","sgizj","qjkqmnefo","g","zaa","epw","fz","ycad","oocqnz","bwowse","bcnealx","imlwxoxzml","crmahpv","vh","qfwyd","dycxopdrzb","hvbje","f","qwowse","lnwgxmqdc","k","fmffskrv","bmlwxoxzml","iffrcpdsu","xpqodjcvri","cmsyrkrrm","unb","ln","hmffskrv","bihbm","fv","eqwg","p","vnpooqdpe","okja","lmffskrv","r","pnwgxmqdc","gpyamphhda","azklwrs","ejwuuqz","rmxhrnzz","gggwj","qh","siantvj","pcad","y","nteqr","mi","ri","jqwg","fnpooqdpe","ukqhjnk","tyaastc","xmsyrkrrm","l","ypmpilxnl","lzf","nkja","xiantvj","fvbje","ooltwf","fycxopdrzb","gi","bmffskrv","bnihnwsx","epci","awowse","tokjknsyk","owan","vz","qd","ng","goebwzhefq","xafrsm","sjyfskc","bfr","nwan","mjkj","ujlr","mrmahpv"], 5808), # pylint: disable=line-too-long
        ]:
            output = self.distinctNames(ideas)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
