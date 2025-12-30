from typing import List, Tuple, Dict
import heapq
from collections import defaultdict
import unittest

# https://leetcode.com/problems/the-skyline-problem/

# python3 -m unittest heaps/0218-the-skyline-problem.py


class Solution(unittest.TestCase):
    # Approach: Sweep Line + Max Heap
    # Create events for each building's start and end
    # Use max heap to track active building heights at each x-coordinate
    # When max height changes, add key point to result
    # N = number of buildings
    # Time: O(N log N) - N events, each heap operation is O(log N)
    # Space: O(N) - heap and events storage
    def getSkyline(self, buildings: List[List[int]]) -> List[List[int]]:
        # Create events: (x, isStart, height)
        # For start: use negative height for sorting (taller buildings first)
        # For end: use positive height
        events: List[Tuple[int, int, int]] = []
        for left, right, height in buildings:
            # Start event: (x, 0, -height) - 0 means start, negative for max heap
            events.append((left, 0, -height))
            # End event: (x, 1, height) - 1 means end
            events.append((right, 1, height))

        # Sort events: by x, then by type (start before end), then by height
        events.sort()

        result: List[List[int]] = []
        # Max heap: use negative heights (Python has min heap)
        max_heap = [0]  # Ground level
        # Track height frequencies (for removal)
        height_count: Dict[int, int] = defaultdict(int)
        height_count[0] = 1

        i = 0
        while i < len(events):
            curr_x = events[i][0]

            # Process all events at the same x-coordinate
            while i < len(events) and events[i][0] == curr_x:
                _, event_type, height = events[i]

                if event_type == 0:  # Start event
                    h = -height  # Convert back to positive
                    heapq.heappush(max_heap, -h)  # Push as negative for max heap
                    height_count[h] += 1
                else:  # End event
                    height_count[height] -= 1
                    if height_count[height] == 0:
                        del height_count[height]

                i += 1

            # Remove invalid heights from top of heap
            while max_heap and -max_heap[0] not in height_count:
                heapq.heappop(max_heap)

            # Get current max height
            max_height = -max_heap[0] if max_heap else 0

            # If height changed, add to result
            if not result or result[-1][1] != max_height:
                result.append([curr_x, max_height])

        return result

    def test(self):
        for buildings, expected in [
            (
                [[2, 9, 10], [3, 7, 15], [5, 12, 12], [15, 20, 10], [19, 24, 8]],
                [[2, 10], [3, 15], [7, 12], [12, 0], [15, 10], [20, 8], [24, 0]]
            ),
            (
                [[0, 2, 3], [2, 5, 3]],
                [[0, 3], [5, 0]]
            ),
            (
                [[0, 3, 3]],
                [[0, 3], [3, 0]]
            ),
            (
                [[1, 2, 1], [1, 2, 2], [1, 2, 3]],
                [[1, 3], [2, 0]]
            ),
            (
                [[0, 2, 3], [0, 2, 3]],
                [[0, 3], [2, 0]]
            ),
        ]:
            output = self.getSkyline(buildings)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
