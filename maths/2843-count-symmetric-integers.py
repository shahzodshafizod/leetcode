import unittest

# https://leetcode.com/problems/count-symmetric-integers/

# python3 -m unittest maths/2843-count-symmetric-integers.py

class Solution(unittest.TestCase):
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        # calculate length of low
        # and the next length's start number
        tmp, length, next = low, 0, 1
        while tmp > 0:
            tmp //= 10
            length += 1
            next *= 10
        count = 0
        num = low
        while num <= high:
            # if number became one digit longer
            # we change length and the next length's start number
            if num == next:
                length += 1
                next *= 10
            # if length is odd, it cannot be split into two parts
            # so we can skip the whole current odd-length section
            if length&1 == 1:
                num = next
                continue
            # now, we add right part digits to total
            total, tmp = 0, num
            for idx in range(length//2):
                total += tmp%10
                tmp //= 10
            # but substract left digits from total
            while tmp > 0:
                total -= tmp%10
                tmp //= 10
            # if left and right parts have equal sums, total should be 0
            if total == 0:
                count += 1
            num += 1
        return count

    def test(self):
        for low, high, expected in [
            (1, 100, 9),
            (1200, 1230, 4),
        ]:
            output = self.countSymmetricIntegers(low, high)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
