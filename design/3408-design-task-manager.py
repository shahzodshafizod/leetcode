import heapq
from collections import defaultdict
from typing import List, Any, Optional, Dict, Tuple
import unittest

# https://leetcode.com/problems/design-task-manager/

# python3 -m unittest design/3408-design-task-manager.py


class TaskManager:

    def __init__(self, tasks: List[List[int]]):
        self.tasks: List[Tuple[int, int]] = []  # [(-priority, -taskId)]
        self.priorities: Dict[int, Tuple[int, int]] = defaultdict(lambda: (-1, 0))  # taskId -> (priority, userId)
        for userId, taskId, priority in tasks:
            heapq.heappush(self.tasks, (-priority, -taskId))
            self.priorities[taskId] = (priority, userId)

    def add(self, userId: int, taskId: int, priority: int) -> None:
        heapq.heappush(self.tasks, (-priority, -taskId))
        self.priorities[taskId] = (priority, userId)

    def edit(self, taskId: int, newPriority: int) -> None:
        heapq.heappush(self.tasks, (-newPriority, -taskId))
        self.priorities[taskId] = (newPriority, self.priorities[taskId][1])

    def rmv(self, taskId: int) -> None:
        self.priorities[taskId] = (-1, self.priorities[taskId][1])

    def execTop(self) -> int:
        while self.tasks and -self.tasks[0][0] != self.priorities[-self.tasks[0][1]][0]:
            heapq.heappop(self.tasks)
        userId = -1
        if self.tasks:
            _, taskId = heapq.heappop(self.tasks)
            userId = self.priorities[-taskId][1]
            self.rmv(-taskId)
        return userId


# Your TaskManager object will be instantiated and called as such:
# obj = TaskManager(tasks)
# obj.add(userId,taskId,priority)
# obj.edit(taskId,newPriority)
# obj.rmv(taskId)
# param_4 = obj.execTop()


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[tuple[List[str], List[Any], List[Any]]] = [
            (
                ["TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"],
                [[[[1, 101, 10], [2, 102, 20], [3, 103, 15]]], [4, 104, 5], [102, 8], [], [101], [5, 105, 15], []],
                [None, None, None, 3, None, None, 5],
            ),
            (
                ["TaskManager", "edit", "execTop", "edit", "rmv", "rmv", "execTop", "execTop"],
                [[[[7, 13, 16], [3, 11, 41], [7, 18, 40]]], [13, 0], [], [13, 19], [18], [13], [], []],
                [None, None, 3, None, None, None, -1, -1],
            ),
        ]
        for commands, values, expected in test_cases:
            manager: Optional[TaskManager] = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "TaskManager":
                        manager = TaskManager(values[idx][0])
                    case "add":
                        if manager is not None:
                            manager.add(values[idx][0], values[idx][1], values[idx][2])
                    case "edit":
                        if manager is not None:
                            manager.edit(values[idx][0], values[idx][1])
                    case "rmv":
                        if manager is not None:
                            manager.rmv(values[idx][0])
                    case "execTop":
                        if manager is not None:
                            output = manager.execTop()
                    case _:
                        pass
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
