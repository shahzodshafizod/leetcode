import unittest

# https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/

# python3 -m unittest maths/3443-maximum-manhattan-distance-after-k-changes.py


class Solution(unittest.TestCase):
    # # Approach #1
    # # Time: O(n)
    # # Space: O(1)
    # def maxDistance(self, s: str, k: int) -> int:
    #     changer = {'N':'S', 'S':'N', 'E':'W', 'W':'E'}
    #     distance = 0
    #     for hor in ['E','W']:
    #         for ver in ['N','S']:
    #             quota, x, y = k, 0, 0
    #             for c in s:
    #                 if quota and (c == hor or c == ver):
    #                     c = changer[c]
    #                     quota -= 1
    #                 x += int(c == 'E') - int(c == 'W')
    #                 y += int(c == 'N') - int(c == 'S')
    #                 distance = max(distance, abs(x)+abs(y))
    #     return distance

    # # Approach #2
    # # Time: O(n)
    # # Space: O(1)
    # def maxDistance(self, s: str, k: int) -> int:
    #     cnt = defaultdict(int)
    #     distance = 0
    #     for c in s:
    #         cnt[c] += 1

    #         vmax, vmin = max(cnt['N'],cnt['S']), min(cnt['N'],cnt['S'])
    #         hmax, hmin = max(cnt['E'],cnt['W']), min(cnt['E'],cnt['W'])

    #         dist = vmax + 2*min(k, vmin) - vmin
    #         dist += hmax + 2*min(k-min(k, vmin), hmin) - hmin

    #         distance = max(distance, dist)
    #     return distance

    # Approach #3
    # Time: O(n)
    # Space: O(1)
    def maxDistance(self, s: str, k: int) -> int:
        distance, lon, lat = 0, 0, 0
        for idx, c in enumerate(s):
            match c:
                case 'E':
                    lon += 1
                case 'W':
                    lon -= 1
                case 'N':
                    lat += 1
                case 'S':
                    lat -= 1
                case _:
                    pass
            distance = max(distance, min(abs(lon) + abs(lat) + 2 * k, idx + 1))
        return distance

    def test(self):
        for s, k, expected in [
            ("NWSE", 1, 3),
            ("NSWWEW", 3, 6),
            ("EWWE", 0, 1),
            ("NSES", 1, 4),
            ("SN", 0, 1),
        ]:
            output = self.maxDistance(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
