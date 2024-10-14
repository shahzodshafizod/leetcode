import heapq
from typing import List
import unittest

# https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

# python3 -m unittest heaps/0632-smallest-range-covering-elements-from-k-lists.py

"""
[4,10,15,24,26] | [0,9,12,20] | [5,18,22,30] = [0,5]
[4,10,15,24,26] | [9,12,20]   | [5,18,22,30] = [4,9]
[10,15,24,26]   | [9,12,20]   | [5,18,22,30] = [5,10]
[10,15,24,26]   | [9,12,20]   | [18,22,30]   = [9,18]
[10,15,24,26]   | [12,20]     | [18,22,30]   = [10,18]
[15,24,26]      | [12,20]     | [18,22,30]   = [12,18]
[24,26]         | [20]        | [18,22,30]   = [18,24]
[24,26]         | [20]        | [22,30]      = [20,24]
[24,26]         | []          | [22,30]      = stop
"""

class Solution(unittest.TestCase):
    # # Approach: Heap (Priority Queue)
    # # Time: O(N+KLogN), N=len(nums), K=count(all nums)
    # # Space: O(N)
    # def smallestRange(self, nums: List[List[int]]) -> List[int]:
    #     n = len(nums)
    #     length = 2 * int(1e5)
    #     pq = [] # (num,row,col), nums[row][col]
    #     end = int(-1e5)
    #     for idx in range(n): # O(N)
    #         heapq.heappush(pq, (nums[idx][0], idx, 0))
    #         end = max(end, nums[idx][0])
    #     while True: # O(min(len(nums[i])))
    #         start, row, col = pq[0]
    #         if end-start < length:
    #             length = end-start
    #             interval = [start,end]
    #         if col+1 == len(nums[row]):
    #             break
    #         col += 1
    #         end = max(end, nums[row][col])
    #         heapq.heapreplace(pq, (nums[row][col], row, col)) # O(LogN)
    #     return interval

    # Approach: Two Pointer
    # Time: O(KLogK), K=count(all nums)
    # Space: O(K)
    def smallestRange(self, nums: List[List[int]]) -> List[int]:
        n = len(nums)
        merged = []
        for idx in range(n):
            for num in nums[idx]:
                merged.append((num,idx))
        merged.sort()
        freq = [0] * n
        min_length = 2*int(1e5)
        left = count = 0
        for right in range(len(merged)):
            freq[merged[right][1]] += 1
            if freq[merged[right][1]] == 1:
                count += 1
            while count == n:
                length = merged[right][0] - merged[left][0]
                if length < min_length:
                    min_length = length
                    interval = [merged[left][0], merged[right][0]]
                freq[merged[left][1]] -= 1
                if freq[merged[left][1]] == 0:
                    count -= 1
                left += 1
        return interval

    def test(self) -> None:
        for nums, expected in [
            ([[4,10,15,24,26],[0,9,12,20],[5,18,22,30]], [20,24]),
            ([[1,2,3],[1,2,3],[1,2,3]], [1,1]),
            ([[10],[11]], [10,11]),
            ([[2,18,24,26],[0,10,12,20],[1,3,22,30]], [0,2]),
            ([[-38,15,17,18],[-34,46,58,59,61],[-55,-31,-13,64,82,82,83,84,85],[-3,63,70,90],[2,6,10,28,28,32,32,32,33],[-23,82,88,88,88,89],[33,60,72,74,75],[-5,44,44,57,58,58,60],[-29,-22,-4,-4,17,18,19,19,19,20],[22,57,82,89,93,94],[24,38,45],[-100,-56,41,49,50,53,53,54],[-76,-69,-66,-53,-27,-1,9,29,31,32,32,32,34],[22,47,56],[-34,-28,7,44]], [18,82]),
            ([[35,77,78,78,79],[34,49,77],[-11,-2,9,14,30,31,40,53,57,60,61,63,65],[22,30],[-28,29,44,54,54,54,56],[-77,22,59,59,66,67],[76,77,77,83,84,85],[-80,14,35,38,42,42,45,48],[19,27,34,73,73,80,80,81],[2,6,20,22,25,28,29],[36,42,45,46,47,47,47,47,47,47,48],[-30,6,35,37,37,38],[20,30,71,71,75,78,79],[37,47,47,47,49],[-53,-8,27,38,46,66,68,69],[-74,15,33,35,48,55],[-14,40,41,56,56,67,69,70,70,74,74,76],[7,10,36,57,78,78,79,81]], [29,76]),
            ([[-1,1],[-2,2],[-3,3],[-4,4],[-5,5],[-6,6],[-7,7],[-8,8],[-9,9],[-10,10],[-11,11],[-12,12],[-13,13],[-14,14],[-15,15],[-16,16],[-17,17],[-18,18],[-19,19],[-20,20],[-21,21],[-22,22],[-23,23],[-24,24],[-25,25],[-26,26],[-27,27],[-28,28],[-29,29],[-30,30],[-31,31],[-32,32],[-33,33],[-34,34],[-35,35],[-36,36],[-37,37],[-38,38],[-39,39],[-40,40],[-41,41],[-42,42],[-43,43],[-44,44],[-45,45],[-46,46],[-47,47],[-48,48],[-49,49],[-50,50],[-51,51],[-52,52],[-53,53],[-54,54],[-55,55]], [-55,-1]),
            ([[-98573,-83574,-38270,36089,89655,96876,96896,98519,99958,99987,99988,100000],[-27200,42327,86435,87214,93353,98904,99296,99497,99966,99986,99999,99999,100000],[90807,99350,99422,99799,99835,99862,99981,99998,100000],[-16594,-15615,9783,79955,91914,99904,99919,99972,99987,99993,99999,99999,100000],[436,4988,12287,16801,22176,46184,80686,98543,99017,99373,99504,99848,99993,99994,99998,100000],[-50242,69111,89964,97763,98549,99739,99927,99934,99945,99984,99985,99990,99993,99998,99999,100000]], [100000,100000]),
            # ([[-91759,75526,81112,83549,98897,99823,99969,99993,99994,99997,100000],[14040,77829,84459,90363,98551,99596,99977,99985,99989,99993,99995,99997,99999,100000],[89452,99027,99534,99936,99964,99965,99991,100000],[78304,98496,99615,99745,99853,99971,99982,99989,99990,99990,100000],[32538,71963,96142,97679,99622,99638,99995,100000],[-17133,41836,76650,98605,98980,99404,99498,99714,99868,99983,99984,99990,99993,99995,99995,99997,99999,99999,99999,99999,100000],[36079,52648,81682,91633,92735,95810,99277,99943,99963,99975,99993,99999,99999,99999,100000],[76800,77599,95574,96000,97714,98680,99552,99805,99831,99858,99866,99912,99971,99982,99985,99989,99993,99993,99997,100000],[5636,42324,56011,98351,98668,99491,99694,99914,99923,99931,99963,99973,99974,99996,100000],[-5694,45326,54082,97544,98217,99383,99571,99596,99713,99729,99856,99903,99947,99957,99996,99997,99997,99997,99997,99997,99997,99997,99998,100000],[-6838,26837,82850,97297,97607,98957,99917,99927,99930,99931,99972,99973,99988,99989,99994,99996,99999,99999,100000],[-10825,48669,85359,91160,98655,98740,98870,99635,99861,99909,99923,99934,99979,99997,99998,99999,100000],[-60721,-58245,50091,78904,88063,92528,96985,98873,99887,99919,99984,99992]], [99990,99997]),
            # ([[95387,95790,97307,98168,99868,99995,99995,100000],[-69454,-17042,8172,50983,63432,72854,73012,80848,83723,85916,91759,99779,99913,99944,99994,99999,99999],[65641,95910,97995,98196,98969,99008,99591,99732,99735,99896,99952,99989,99999,100000],[57459,95855,97360,98320,99147,99865,99955,99989,99997,99998,100000],[-81589,-3474,84141,92981,95255,99192,99962,99970,99994,99998,99999,100000],[-23262,92924,95548,96462,99338,99553,99555,99569,99644,99903,99909,99999,99999,100000],[-58466,24432,87898,92795,95701,98143,98163,99182,99351,99746,99811,99943,99955,99978,99997,100000],[-97588,7867,10356,20288,67836,69868,73038,77753,81937,88474,89979,92182,98091,99635,99902,99941,99975,99987,99991,99998,99998,99998,99998,99998,99999,99999,99999,100000],[-96955,41521,84537,89794,96226,97103,97490,99347,99957,99997,100000],[-49247,93963,99006,99428,99964,99992,100000],[46062,48599,95745,98620,98677,99516,99802,99973,99993,100000],[-3786,59724,62870,80033,90471,98836,99395,99574,99682,99724,99909,99963,99979,99999,100000],[-62512,-19463,84187,89388,91368,95524,98987,99085,99230,99809,99978,100000],[18183,83019,98718,99570,99777,99980,100000],[19925,20448,81509,93698,98451,98776,98915,99007,99925,99994,99996,99999,100000],[-96129,93245,95417,98492,99013,99921,99934,99989,99995,100000],[-25468,61948,68372,85478,91239,98906,98988,99653,99915,99957,99998,99999,99999,100000],[36648,65266,95679,98905,99868,99977,99983,99983,99995,99995,99996,99997,100000],[56006,78969,86785,89834,92494,93887,98268,99771,99982,99998,99999,100000],[95387,95790,97307,98168,99868,99995,99995,100000],[-69454,-17042,8172,50983,63432,72854,73012,80848,83723,85916,91759,99779,99913,99944,99994,99999,99999],[65641,95910,97995,98196,98969,99008,99591,99732,99735,99896,99952,99989,99999,100000],[57459,95855,97360,98320,99147,99865,99955,99989,99997,99998,100000],[-81589,-3474,84141,92981,95255,99192,99962,99970,99994,99998,99999,100000],[-23262,92924,95548,96462,99338,99553,99555,99569,99644,99903,99909,99999,99999,100000],[-58466,24432,87898,92795,95701,98143,98163,99182,99351,99746,99811,99943,99955,99978,99997,100000],[-97588,7867,10356,20288,67836,69868,73038,77753,81937,88474,89979,92182,98091,99635,99902,99941,99975,99987,99991,99998,99998,99998,99998,99998,99999,99999,99999,100000],[-96955,41521,84537,89794,96226,97103,97490,99347,99957,99997,100000],[-49247,93963,99006,99428,99964,99992,100000],[46062,48599,95745,98620,98677,99516,99802,99973,99993,100000],[-3786,59724,62870,80033,90471,98836,99395,99574,99682,99724,99909,99963,99979,99999,100000],[-62512,-19463,84187,89388,91368,95524,98987,99085,99230,99809,99978,100000],[18183,83019,98718,99570,99777,99980,100000],[19925,20448,81509,93698,98451,98776,98915,99007,99925,99994,99996,99999,100000],[-96129,93245,95417,98492,99013,99921,99934,99989,99995,100000],[-25468,61948,68372,85478,91239,98906,98988,99653,99915,99957,99998,99999,99999,100000],[36648,65266,95679,98905,99868,99977,99983,99983,99995,99995,99996,99997,100000],[56006,78969,86785,89834,92494,93887,98268,99771,99982,99998,99999,100000],[95387,95790,97307,98168,99868,99995,99995,100000],[-69454,-17042,8172,50983,63432,72854,73012,80848,83723,85916,91759,99779,99913,99944,99994,99999,99999],[65641,95910,97995,98196,98969,99008,99591,99732,99735,99896,99952,99989,99999,100000],[57459,95855,97360,98320,99147,99865,99955,99989,99997,99998,100000],[-81589,-3474,84141,92981,95255,99192,99962,99970,99994,99998,99999,100000],[-23262,92924,95548,96462,99338,99553,99555,99569,99644,99903,99909,99999,99999,100000],[-58466,24432,87898,92795,95701,98143,98163,99182,99351,99746,99811,99943,99955,99978,99997,100000],[-97588,7867,10356,20288,67836,69868,73038,77753,81937,88474,89979,92182,98091,99635,99902,99941,99975,99987,99991,99998,99998,99998,99998,99998,99999,99999,99999,100000],[-96955,41521,84537,89794,96226,97103,97490,99347,99957,99997,100000],[-49247,93963,99006,99428,99964,99992,100000],[46062,48599,95745,98620,98677,99516,99802,99973,99993,100000],[-3786,59724,62870,80033,90471,98836,99395,99574,99682,99724,99909,99963,99979,99999,100000],[-62512,-19463,84187,89388,91368,95524,98987,99085,99230,99809,99978,100000],[18183,83019,98718,99570,99777,99980,100000],[19925,20448,81509,93698,98451,98776,98915,99007,99925,99994,99996,99999,100000],[-96129,93245,95417,98492,99013,99921,99934,99989,99995,100000],[-25468,61948,68372,85478,91239,98906,98988,99653,99915,99957,99998,99999,99999,100000],[36648,65266,95679,98905,99868,99977,99983,99983,99995,99995,99996,99997,100000],[56006,78969,86785,89834,92494,93887,98268,99771,99982,99998,99999,100000],[95387,95790,97307,98168,99868,99995,99995,100000],[-69454,-17042,8172,50983,63432,72854,73012,80848,83723,85916,91759,99779,99913,99944,99994,99999,99999],[65641,95910,97995,98196,98969,99008,99591,99732,99735,99896,99952,99989,99999,100000],[57459,95855,97360,98320,99147,99865,99955,99989,99997,99998,100000],[-81589,-3474,84141,92981,95255,99192,99962,99970,99994,99998,99999,100000],[-23262,92924,95548,96462,99338,99553,99555,99569,99644,99903,99909,99999,99999,100000],[-58466,24432,87898,92795,95701,98143,98163,99182,99351,99746,99811,99943,99955,99978,99997,100000],[-97588,7867,10356,20288,67836,69868,73038,77753,81937,88474,89979,92182,98091,99635,99902,99941,99975,99987,99991,99998,99998,99998,99998,99998,99999,99999,99999,100000],[-96955,41521,84537,89794,96226,97103,97490,99347,99957,99997,100000],[-49247,93963,99006,99428,99964,99992,100000],[46062,48599,95745,98620,98677,99516,99802,99973,99993,100000],[-3786,59724,62870,80033,90471,98836,99395,99574,99682,99724,99909,99963,99979,99999,100000],[-62512,-19463,84187,89388,91368,95524,98987,99085,99230,99809,99978,100000],[18183,83019,98718,99570,99777,99980,100000],[19925,20448,81509,93698,98451,98776,98915,99007,99925,99994,99996,99999,100000],[-96129,93245,95417,98492,99013,99921,99934,99989,99995,100000],[-25468,61948,68372,85478,91239,98906,98988,99653,99915,99957,99998,99999,99999,100000],[36648,65266,95679,98905,99868,99977,99983,99983,99995,99995,99996,99997,100000],[56006,78969,86785,89834,92494,93887,98268,99771,99982,99998,99999,100000],[95387,95790,97307,98168,99868,99995,99995,100000],[-69454,-17042,8172,50983,63432,72854,73012,80848,83723,85916,91759,99779,99913,99944,99994,99999,99999],[65641,95910,97995,98196,98969,99008,99591,99732,99735,99896,99952,99989,99999,100000],[57459,95855,97360,98320,99147,99865,99955,99989,99997,99998,100000],[-81589,-3474,84141,92981,95255,99192,99962,99970,99994,99998,99999,100000],[-23262,92924,95548,96462,99338,99553,99555,99569,99644,99903,99909,99999,99999,100000],[-58466,24432,87898,92795,95701,98143,98163,99182,99351,99746,99811,99943,99955,99978,99997,100000],[-97588,7867,10356,20288,67836,69868,73038,77753,81937,88474,89979,92182,98091,99635,99902,99941,99975,99987,99991,99998,99998,99998,99998,99998,99999,99999,99999,100000],[-96955,41521,84537,89794,96226,97103,97490,99347,99957,99997,100000],[-49247,93963,99006,99428,99964,99992,100000],[46062,48599,95745,98620,98677,99516,99802,99973,99993,100000],[-3786,59724,62870,80033,90471,98836,99395,99574,99682,99724,99909,99963,99979,99999,100000],[-62512,-19463,84187,89388,91368,95524,98987,99085,99230,99809,99978,100000],[18183,83019,98718,99570,99777,99980,100000],[19925,20448,81509,93698,98451,98776,98915,99007,99925,99994,99996,99999,100000],[-96129,93245,95417,98492,99013,99921,99934,99989,99995,100000],[-25468,61948,68372,85478,91239,98906,98988,99653,99915,99957,99998,99999,99999,100000],[36648,65266,95679,98905,99868,99977,99983,99983,99995,99995,99996,99997,100000],[56006,78969,86785,89834,92494,93887,98268,99771,99982,99998,99999,100000],[95387,95790,97307,98168,99868,99995,99995,100000],[-69454,-17042,8172,50983,63432,72854,73012,80848,83723,85916,91759,99779,99913,99944,99994,99999,99999],[65641,95910,97995,98196,98969,99008,99591,99732,99735,99896,99952,99989,99999,100000],[57459,95855,97360,98320,99147,99865,99955,99989,99997,99998,100000],[-81589,-3474,84141,92981,95255,99192,99962,99970,99994,99998,99999,100000],[-23262,92924,95548,96462,99338,99553,99555,99569,99644,99903,99909,99999,99999,100000],[-58466,24432,87898,92795,95701,98143,98163,99182,99351,99746,99811,99943,99955,99978,99997,100000],[-97588,7867,10356,20288,67836,69868,73038,77753,81937,88474,89979,92182,98091,99635,99902,99941,99975,99987,99991,99998,99998,99998,99998,99998,99999,99999,99999,100000],[-96955,41521,84537,89794,96226,97103,97490,99347,99957,99997,100000],[-49247,93963,99006,99428,99964,99992,100000],[46062,48599,95745,98620,98677,99516,99802,99973,99993,100000],[-3786,59724,62870,80033,90471,98836,99395,99574,99682,99724,99909,99963,99979,99999,100000],[-62512,-19463,84187,89388,91368,95524,98987,99085,99230,99809,99978,100000],[18183,83019,98718,99570,99777,99980,100000],[19925,20448,81509,93698,98451,98776,98915,99007,99925,99994,99996,99999,100000],[-96129,93245,95417,98492,99013,99921,99934,99989,99995,100000],[-25468,61948,68372,85478,91239,98906,98988,99653,99915,99957,99998,99999,99999,100000],[36648,65266,95679,98905,99868,99977,99983,99983,99995,99995,99996,99997,100000],[56006,78969,86785,89834,92494,93887,98268,99771,99982,99998,99999,100000]], [99999,100000]),
        ]:
            output = self.smallestRange(nums)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
