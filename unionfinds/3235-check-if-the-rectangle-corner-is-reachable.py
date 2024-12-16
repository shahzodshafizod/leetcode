from typing import List
import unittest

# https://leetcode.com/problems/check-if-the-rectangle-corner-is-reachable/

# python3 -m unittest unionfinds/3235-check-if-the-rectangle-corner-is-reachable.py

# TODO: Solve this problem by yourself

class Solution(unittest.TestCase):
    # # Approach #1: Graph
    # # Time: O(N^2)
    # # Space: O(N)
    # def canReachCorner(self, xCorner: int, yCorner: int, circles: List[List[int]]) -> bool:
    #     # inCircle returns whether Point (px, py) is located
    #     # inside Circle (cx, cy, r) (including the border)
    #     def inCircle(circle: List[int], px: int, py: int) -> bool:
    #         cx, cy, radius = circle
    #         return (cx-px)*(cx-px)+(cy-py)*(cy-py) <= radius*radius
    #     visited = [False] * len(circles)
    #     def dfs(i: int) -> bool:
    #         x1, y1, r1 = circles[i]
    #         # Circle i intersects with the bottom or right border of the rectangle
    #         if y1 <= yCorner and abs(x1-xCorner) <= r1 or \
    #             x1 <= xCorner and y1 <= r1 or \
    #             x1 > xCorner and inCircle(circles[i], xCorner, 0):
    #             return True
    #         visited[i] = True
    #         for j, circle in enumerate(circles):
    #             if visited[j]:
    #                 continue
    #             x2, y2, r2 = circle
    #             # Let Point A which |O1A| / |O1O2| = r1 / (r1+r2).
    #             # If two circles are connected to each other then A must be inside the intersection
    #             # And its coordinate is: (x1·r2+x2·r1)/(r1+r2), (y1·r2+y2·r1)/(r1+r2)
    #             distance = (x1-x2)**2 + (y1-y2)**2
    #             rdistance = (r1 + r2) ** 2
    #             if distance <= rdistance and \
    #                 x1*r2+x2*r1 < (r1+r2)*xCorner and \
    #                 y1*r2+y2*r1 < (r1+r2)*yCorner and \
    #                 dfs(j):
    #                 return True
    #         return False
    #     for i, circle in enumerate(circles):
    #         x, y, r = circle
    #         # DFS starts from circles which intersects the left/top border of the rectangle
    #         if inCircle(circle, 0, 0):
    #             return False
    #         if inCircle(circle, xCorner, yCorner):
    #             return False
    #         if visited[i]:
    #             continue
    #         if (x <= xCorner and abs(y-yCorner) <= r or \
    #             y <= yCorner and x <= r or \
    #             y > yCorner and inCircle(circle, 0, yCorner)) and \
    #             dfs(i):
    #             return False
    #     return True

    # # Approach #2: Graph
    # # Time: O(N^2)
    # # Space: O(N)
    # def canReachCorner(self, xCorner: int, yCorner: int, circles: List[List[int]]) -> bool:
    #     n = len(circles)
    #     edges = [0] * n
    #     inside = [False] * n
    #     for idx in range(n):
    #         x, y, r = circles[idx]
    #         inside[idx] = x <= xCorner and y <= yCorner
    #         # if the circle intersects with start/end point
    #         if x*x+y*y <= r*r or (x-xCorner)*(x-xCorner)+(y-yCorner)*(y-yCorner) <= r*r:
    #             return False
    #         # check horizontal edges
    #         if y-r <= 0 and y+r >= 0 and x <= xCorner or x+r >= xCorner and x-r <= xCorner and y <= yCorner:
    #             edges[idx] = 1
    #         # check vertical edges
    #         if x-r <= 0 and x+r >= 0 and y <= yCorner or y+r >= yCorner and y-r <= yCorner and x <= xCorner:
    #             edges[idx] += 2
    #         if edges[idx] == 3:
    #             return False
    #     seen = [False] * n
    #     def dfs(idx: int) -> int:
    #         seen[idx] = True
    #         mask = edges[idx]
    #         x, y, r = circles[idx]
    #         for j in range(n):
    #             if seen[j]: continue
    #             x2, y2, r2 = circles[j]
    #             dx = x - x2
    #             dy = y - y2
    #             dr = r + r2
    #             if dx*dx+dy*dy <= dr*dr and (inside[idx] or inside[j] or (x+x2 <= 2*xCorner and y+y2 <= 2*yCorner)):
    #                 mask |= dfs(j)
    #         return mask
    #     for idx in range(n):
    #         if not seen[idx] and dfs(idx) == 3:
    #             return False
    #     return True

    # Approach: Union-Find
    # Time: O(N)
    # Space: O(N)
    def canReachCorner(self, xCorner: int, yCorner: int, circles: List[List[int]]) -> bool:
        n = len(circles)
        parent = list(range(n+2))
        def find(x: int) -> int:
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]
        for idx in range(n):
            x, y, r = circles[idx]
            # if the circle overlaps right-top
            if (x-xCorner)*(x-xCorner)+(y-yCorner)*(y-yCorner) <= r*r:
                return False
            # # if the circle overlaps left-bottom
            # if x*x+y*y <= r*r:
            # 	return False
            if x <= r and y <= yCorner or y+r >= yCorner and x <= xCorner:
                parent[find(n)] = find(idx)
            if y <= r and x <= xCorner or x+r >= xCorner and y <= yCorner:
                parent[find(n+1)] = find(idx)
            # if the circle touches both borders
            if find(n) == find(n+1):
                return False
        # if either border is untouched
        if find(n) == n or find(n+1) == n+1:
            return True
        for i in range(n):
            x, y, r = circles[i]
            # if circle 1 is out of usable range
            if x-r >= xCorner or y-r >= yCorner or x >= xCorner and y >= yCorner:
                continue
            for j in range(i):
                # if already unioned
                if find(i) == find(j):
                    continue
                x2, y2, r2 = circles[j]
                # if pair is out of usable range
                if (x+x2)/2 >= xCorner and (y+y2)/2 >= yCorner:
                    continue
                # if circles intersect
                if (x-x2)*(x-x2)+(y-y2)*(y-y2) <= (r+r2)*(r+r2):
                    parent[find(i)] = find(j)
                    # if both borders are unioned
                    if find(n) == find(n+1):
                        return False
        return True

    def test(self):
        for xCorner, yCorner, circles, expected in [
            (3, 4, [[2,1,1]], True),
            (3, 3, [[1,1,2]], False),
            (3, 3, [[2,1,1],[1,2,1]], False),
            (4, 4, [[5,5,1]], True),
            (3, 3, [[2,1000,997],[1000,2,997]], True),
            (3, 4, [[5,6,2]], True),
            (3, 3, [[1,100,2],[3,100,2]], True),
            (4, 4, [[2,6,2],[6,2,2],[6,6,2]], True),
            (3, 3, [[4,4,2]], False),
            (22742157, 210809967, [[22741186,210810964,200],[22741869,210809432,165],[22742256,210810275,182],[22742089,210809693,129],[22741912,210810128,196],[22741658,210809205,144],[22741648,210809094,118],[22741920,210809808,128]], False),
            (8, 10, [[5,2,2],[6,7,2],[3,1,1],[7,5,1],[5,3,1],[3,7,3],[1,7,1]], False),
        ]:
            output = self.canReachCorner(xCorner, yCorner, circles)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
