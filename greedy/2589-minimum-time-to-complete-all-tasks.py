import unittest

# https://leetcode.com/problems/minimum-time-to-complete-all-tasks/

# python3 -m unittest greedy/2589-minimum-time-to-complete-all-tasks.py


class Solution(unittest.TestCase):
    # Approach: Greedy Algorithm with Sorting and Interval Scheduling
    # Time Complexity: O(n*log(n) + n*m) where n=tasks, m=max(end)
    # Space Complexity: O(m) where m=max(end)
    def findMinimumTime(self, tasks: list[list[int]]) -> int:
        tasks.sort(key=lambda x: x[1])

        total = 0
        on = [False] * 2001

        for start, end, duration in tasks:
            covered = sum(1 for t in range(start, end + 1) if on[t])
            duration -= covered

            time = end
            while time >= start and duration > 0:
                if not on[time]:
                    on[time] = True
                    total += 1
                    duration -= 1
                time -= 1

        return total

    def test(self):
        for tasks, expected in [
            ([[2, 3, 1], [4, 5, 1], [1, 5, 2]], 2),
            ([[1, 3, 2], [2, 5, 3], [5, 6, 2]], 4),
            ([[2, 3, 1], [4, 5, 1], [1, 5, 2]], 2),
            ([[1, 3, 2], [2, 5, 3], [5, 6, 2]], 4),
            ([[1, 1, 1]], 1),
            ([[1, 10, 1], [2, 10, 1], [3, 10, 1]], 1),  # All can use time 10
            ([[1, 5, 5]], 5),  # Must use all times 1,2,3,4,5
            ([[1, 4, 3], [2, 4, 2], [3, 4, 1]], 3),  # Times 4,3,2 cover all
            ([[1, 2, 1], [2, 3, 1], [3, 4, 1]], 2),  # Times 2,3 can cover first two
        ]:
            output = self.findMinimumTime(tasks)
            self.assertEqual(expected, output, f"tasks: {tasks}, expected: {expected}, output: {output}")
