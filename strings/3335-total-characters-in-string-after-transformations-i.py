import unittest

# https://leetcode.com/problems/total-characters-in-string-after-transformations-i/

# python3 -m unittest strings/3335-total-characters-in-string-after-transformations-i.py

class Solution(unittest.TestCase):
    def lengthAfterTransformations(self, s: str, t: int) -> int:
        dp = [0] * 26
        for c in s:
            dp[ord(c)-ord('a')] += 1
        for _ in range(t):
            count_z = dp[-1]
            for i in range(25,0,-1):
                dp[i] = dp[i-1]
            dp[0] = count_z
            dp[1] += count_z
        return sum(dp) % int(1e9+7)

    def test(self):
        for s, t, expected in [
            ("z", 100, 16),
            ("azbk", 1, 5),
            ("abcyy", 2, 7),
            ("zwzzugozhwsycuocakus", 1, 24),
            # ("jqktcurgdvlibczdsvnsg", 7517, 79033769),
            # ("hrzmawnweztcskakojfahyvnoctsctwsbagyqmmoheldlpzctduxmhfcwqcbvovoyswjtdzvsheoofocknqddfsjwxfuuhvznxry", 1000, 652815408),
            # ("xdzbhxqcmhezajdhljzsgshikospdeyxrnwylcvcuvfppquqqxcfbvmdlwbzkxjkwzvoyvmpnlxuyulexoqgayvxlvofyjhmxshfprpbhjywofbqhhufezuyccasrodkzmxkwzfhcfxhlrpidoklhgidflvyppajzgecuhumjyglgzqzcusdniuqgylpxlhkknwbwehtaabnioerjnpxxjqhmxftsoukxbfkndssniyhwqfmtcoerxrkkdjepiyrhmnepuwaunubwrixahwaoecretfzbqavlhzavdherptjpkhqrkpopdheswffikuxrvqohccbyphcrirhhjddjqihwxlszdehalyoqqzsimaaxepttwbfpbtqgwhidvzoegkjeqdhndszmrtgloqwerpdsvqhdvmfqxwdmkocqcltqiojgpstzainiukaejurbvuvjbyyodruvuliahiscdmrjmnthehnmhovjwakenmuwbxqeoox", 100000, 571676932),
        ]:
            output = self.lengthAfterTransformations(s, t)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
