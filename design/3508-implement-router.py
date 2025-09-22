from bisect import bisect_left, bisect_right
from collections import deque, defaultdict
from typing import List, Any, Optional, Deque, Set, Dict
import unittest

# https://leetcode.com/problems/implement-router/

# python3 -m unittest design/3508-implement-router.py


class Router:

    def __init__(self, memoryLimit: int):
        self.cap = memoryLimit
        self.packets: Deque[tuple[int, int, int]] = deque()  # (source, destination, timestamp)
        self.unique: Set[tuple[int, int, int]] = set()  # {(source, destination, timestamp)}
        self.dests: Dict[int, List[int]] = defaultdict(list)  # {destination: [timestamp]}
        self.offset: Dict[int, int] = defaultdict(int)

    def addPacket(self, source: int, destination: int, timestamp: int) -> bool:
        key = (source, destination, timestamp)
        if key in self.unique:
            return False
        if len(self.packets) == self.cap:
            self.forwardPacket()
        self.packets.append(key)
        self.unique.add(key)
        self.dests[destination].append(timestamp)
        return True

    def forwardPacket(self) -> List[int]:
        if not self.packets:
            return []
        key = self.packets.popleft()
        self.unique.remove(key)
        self.offset[key[1]] += 1
        return list(key)

    def getCount(self, destination: int, startTime: int, endTime: int) -> int:
        lo = self.offset[destination]
        left = bisect_left(self.dests[destination], startTime, lo)
        right = bisect_right(self.dests[destination], endTime, lo)
        return right - left


# Your Router object will be instantiated and called as such:
# obj = Router(memoryLimit)
# param_1 = obj.addPacket(source,destination,timestamp)
# param_2 = obj.forwardPacket()
# param_3 = obj.getCount(destination,startTime,endTime)


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[tuple[List[str], List[Any], List[Any]]] = [
            (
                ["Router", "addPacket", "addPacket", "addPacket", "addPacket", "addPacket", "forwardPacket", "addPacket", "getCount"],
                [[3], [1, 4, 90], [2, 5, 90], [1, 4, 90], [3, 5, 95], [4, 5, 105], [], [5, 2, 110], [5, 100, 110]],
                [None, True, True, False, True, True, [2, 5, 90], True, 1],
            ),
            (
                ["Router", "addPacket", "forwardPacket", "forwardPacket"],
                [[2], [7, 4, 90], [], []],
                [None, True, [7, 4, 90], []],
            ),
            (
                ["Router", "addPacket", "addPacket", "addPacket", "getCount", "getCount", "addPacket", "forwardPacket", "addPacket"],
                [[5], [2, 3, 1], [5, 2, 5], [2, 3, 5], [3, 4, 4], [3, 5, 5], [3, 2, 5], [], [2, 3, 5]],
                [None, True, True, True, 0, 1, True, [2, 3, 1], False],
            ),
        ]
        for commands, values, expected in test_cases:
            router: Optional[Router] = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "Router":
                        router = Router(values[idx][0])
                    case "addPacket":
                        if router is not None:
                            output = router.addPacket(values[idx][0], values[idx][1], values[idx][2])
                    case "forwardPacket":
                        if router is not None:
                            output = router.forwardPacket()
                    case "getCount":
                        if router is not None:
                            output = router.getCount(values[idx][0], values[idx][1], values[idx][2])
                    case _:
                        pass
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
