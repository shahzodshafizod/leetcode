from typing import List
import unittest

# https://leetcode.com/problems/flipping-an-image/

# python3 -m unittest twopointers/0832-flipping-an-image.py


class Solution(unittest.TestCase):
    def flipAndInvertImage(self, image: List[List[int]]) -> List[List[int]]:
        n = len(image)
        for row in image:
            left, right = 0, n - 1
            while left <= right:
                row[left], row[right] = row[right] ^ 1, row[left] ^ 1
                left += 1
                right -= 1
        return image

    def test(self):
        for image, expected in [
            ([[1, 1, 0], [1, 0, 1], [0, 0, 0]], [[1, 0, 0], [0, 1, 0], [1, 1, 1]]),
            (
                [[1, 1, 0, 0], [1, 0, 0, 1], [0, 1, 1, 1], [1, 0, 1, 0]],
                [[1, 1, 0, 0], [0, 1, 1, 0], [0, 0, 0, 1], [1, 0, 1, 0]],
            ),
        ]:
            output = self.flipAndInvertImage(image)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
