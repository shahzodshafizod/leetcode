from typing import List
import unittest

# https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/

# python3 -m unittest binarysearch/2071-maximum-number-of-tasks-you-can-assign.py


class Solution(unittest.TestCase):
    # Approach: Binary Search + Greedy Selection of Workers
    # Time: O(nlogn + mlogm + klogk), k=Min(n,m)
    # Space: O(n + m)
    def maxTaskAssign(self, tasks: List[int], workers: List[int], pills: int, strength: int) -> int:
        n, m = len(tasks), len(workers)
        tasks.sort()  # O(n log n)
        workers.sort()  # O(m log m)
        deque = [0] * m

        def greedy(mid: int) -> bool:  # O(min(m,n))
            p = pills
            first, last = m - 1, m - 1
            wid = m - 1
            # enumerate each task from largest to smallest
            for i in range(mid - 1, -1, -1):
                while wid >= m - mid and workers[wid] + strength >= tasks[i]:
                    deque[first] = workers[wid]
                    first -= 1
                    wid -= 1
                if first == last:  # not enough strong workers, even with pills
                    return False
                # if the strongest worker can manage the task without a pill, use that worker
                if deque[last] >= tasks[i]:
                    last -= 1
                elif p > 0:
                    first += 1
                    p -= 1
                else:  # not enough pills
                    return False
            return True

        left, right = 0, min(m, n)  # left=0 is important
        while left < right:  # O(log min(m,n))
            mid = (left + right + 1) // 2
            if greedy(mid):
                left = mid
            else:
                right = mid - 1
        return left

    def test(self):
        for tasks, workers, pills, strength, expected in [
            ([3, 2, 1], [0, 3, 3], 1, 1, 3),
            ([5, 4], [0, 0, 0], 1, 5, 1),
            ([10, 15, 30], [0, 10, 10, 10, 10], 3, 10, 2),
            ([2, 5], [1, 3, 4], 1, 1, 2),
            ([5, 9, 8, 5, 9], [1, 6, 4, 2, 6], 1, 5, 3),
            ([10, 10, 10], [0, 0, 0], 3, 10, 3),
            ([8, 10], [4], 0, 6, 0),
        ]:
            output = self.maxTaskAssign(tasks, workers, pills, strength)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
