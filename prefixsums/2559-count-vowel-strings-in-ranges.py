from typing import List
import unittest

# https://leetcode.com/problems/count-vowel-strings-in-ranges/

# python3 -m unittest prefixsums/2559-count-vowel-strings-in-ranges.py

class Solution(unittest.TestCase):
    # Approach: Prefix Sum
    # Time: O(W+Q), W=len(words), Q=len(queries)
    # Space: O(W)
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        vowels = set(['a','e','i','o','u'])
        precount = [0] * (len(words) + 1)
        for idx, w in enumerate(words):
            precount[idx+1] = precount[idx]
            if w[0] in vowels and w[-1] in vowels:
                precount[idx+1] += 1
        ans = [0] * len(queries)
        for idx, (left, right) in enumerate(queries):
            ans[idx] = precount[right+1] - precount[left]
        return ans

    def test(self):
        for words, queries, expected in [
            (["e"], [[0,0]], [1]),
            (["ae","yi"], [[0,1]], [1]),
            (["a","e","i"], [[0,2],[0,1],[2,2]], [3,2,1]),
            (["aba","bcb","ece","aa","e"], [[0,2],[1,4],[1,1]], [2,3,0]),
            # (["ht","kg","ot","ws","ap","kp","ld","fw","bp","uo","rp","wt","yi","ak","ea","au"], [[15,15],[8,12],[10,12],[14,14],[6,10],[15,15],[13,14],[4,12],[8,8],[14,14],[12,12]], [1,1,0,1,1,1,1,1,0,1,0]),
            # (["norugjhgwqxrdrihwtwdfgi","ykywlzmiplhhvwpbvimczuzecl","rqindmyp","pkzwttxpyrjfyhmlhehjd"], [[1,2],[1,2],[3,3],[0,1],[0,0],[2,3],[0,2],[2,2],[1,1],[2,2],[0,2],[0,1],[2,3],[2,3],[2,2],[1,1],[2,2],[3,3],[3,3],[1,2],[1,3],[1,3],[3,3],[0,0],[0,3],[3,3],[3,3],[3,3],[3,3],[0,2],[2,3]], [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]),
            # (["qysocue","axnisoxgsdkujhjjadtsqddmmda","vwervizcudgedrguu","uzoaikzkhuxbzszqara","rzvuejyncmwpbhs","gap","ekpvxaowewrpzzjdukewfr","tawetkkelzvtqfbyxaxtceegx","ismdpgsjnmqyfhrkaaoyhvarkno"], [[4,8],[7,8],[2,8],[0,5],[3,5],[6,6],[3,3],[2,8],[7,7],[7,8],[7,7],[7,8],[7,8],[8,8],[0,5],[8,8],[6,8],[1,5],[2,6],[2,3],[1,6],[3,3],[0,7],[7,7],[8,8],[2,7],[7,7],[4,4],[4,6],[2,5],[5,5],[5,5],[1,2],[4,6],[5,6]], [1,1,2,2,1,0,1,2,0,1,0,1,1,1,2,1,1,2,1,1,2,1,2,0,1,1,0,0,0,1,0,0,1,0,0]),
            # (["a","rmivyakd","kddwnexxssssnvrske","vceguisunlxtldqenxiyfupvnsxdubcnaucpoi","nzwdiataxfkbikbtsjvcbjxtr","wlelgybcaakrxiutsmwnkuyanvcjczenuyaiy","eueryyiayq","bghegfwmwdoayakuzavnaucpur","ukorsxjfkdojcxgjxgmxbghno","pmgbiuzcwbsakwkyspeikpzhnyiqtqtfyephqhl","gsjdpelkbsruooeffnvjwtsidzw","ugeqzndjtogxjkmhkkczdpqzwcu","ppngtecadjsirj","rvfeoxunxaqezkrlr","adkxoxycpinlmcvmq","gfjhpxlzmokcmvhjcrbrpfakspscmju","rgmzhaj","ychktzwdhfuruhpvdjwfsqjhztshcxdey","yifrzmmyzvfk","mircixfzzobcficujgbj","d","pxcmwnqknyfkmafzbyajjildngccadudfziknos","dxmlikjoivggmyasaktllgmfhqpyznc","yqdbiiqexkemebyuitve"], [[0,0],[5,21],[17,22],[19,23],[13,15],[20,23],[21,23],[6,20],[1,8],[15,20],[17,22],[6,6],[1,2],[4,11],[14,23],[7,10],[16,22],[20,22],[21,22],[15,18],[5,16],[17,23]], [1,2,0,0,0,0,0,2,1,0,0,0,0,2,0,1,0,0,0,0,2,0]),
            # (["e","z","sg","gfiufm","wnyvvgqucyjyrcsptzjvawvjfzfnlajg","ctnydlvnhwynaiwss","jxiimicjpevzuulpeorvhfoizdumoce","bk","ezpbrhqni","qlalpzdvcvsiggfzhbvmbnderrrhpxercsyr","prunwbhfusmtd","dqtpiczphvmmebkjymee","yliexnoxitlmipdpgq","hvbmmyuos","duksgf","suwqmrauuuacqqvznlbtqopzzphekllpwjb","jrxdyvtebhmfscsdwovkosnk","cic","awqdbvynrjlfmuehgwwtxskhihicdgiycxwmbm","stipidhjxuktfsairr"], [[11, 11],[18, 18],[19, 19],[8, 9],[13, 14],[15, 17],[9, 19],[18, 19],[8, 14],[10, 17],[10, 14],[12, 16],[2, 12],[17, 17],[2, 3],[6, 6],[16, 18],[9, 14],[14, 19],[7, 17],[19, 19],[17, 18],[8, 11],[19, 19],[13, 16],[11, 16],[16, 18],[19, 19],[13, 14],[7, 16],[9, 19],[6, 9],[18, 18],[5, 16],[10, 16],[7, 9],[18, 19],[6, 13],[10, 18],[6, 17],[6, 17],[18, 18],[6, 17],[5, 12],[2, 11],[11, 17],[9, 10]], [0,0,0,1,0,0,0,0,1,0,0,0,1,0,0,0,0,0,0,1,0,0,1,0,0,0,0,0,0,1,0,1,0,1,0,1,0,1,0,1,1,0,1,1,1,0,0]),
        ]:
            output = self.vowelStrings(words, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
