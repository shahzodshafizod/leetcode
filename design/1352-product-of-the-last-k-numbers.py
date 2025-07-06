import unittest

# https://leetcode.com/problems/product-of-the-last-k-numbers/

# python3 -m unittest design/1352-product-of-the-last-k-numbers.py


class ProductOfNumbers:

    def __init__(self):
        self.presum = [1]

    def add(self, num: int) -> None:
        if num == 0:
            self.presum = [1]
        else:
            self.presum.append(self.presum[-1] * num)

    def getProduct(self, k: int) -> int:
        if k > len(self.presum) - 1:
            return 0
        return self.presum[-1] // self.presum[-1 - k]


# Your ProductOfNumbers object will be instantiated and called as such:
# obj = ProductOfNumbers()
# obj.add(num)
# param_2 = obj.getProduct(k)


class TestProductOfNumbers(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                [
                    "ProductOfNumbers",
                    "add",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "getProduct",
                    "getProduct",
                    "add",
                    "getProduct",
                ],
                [[], [3], [0], [2], [5], [4], [2], [3], [4], [8], [2]],
                [None, None, None, None, None, None, 20, 40, 0, None, 32],
            ),
            (
                [
                    "ProductOfNumbers",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "add",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "add",
                    "getProduct",
                    "add",
                    "getProduct",
                    "getProduct",
                    "add",
                    "add",
                    "add",
                    "add",
                ],
                [
                    [],
                    [2],
                    [2],
                    [5],
                    [10],
                    [2],
                    [0],
                    [5],
                    [0],
                    [2],
                    [3],
                    [3],
                    [10],
                    [4],
                    [7],
                    [2],
                    [1],
                    [8],
                    [2],
                    [6],
                    [5],
                ],
                [
                    None,
                    None,
                    None,
                    None,
                    None,
                    50,
                    None,
                    None,
                    None,
                    None,
                    None,
                    0,
                    None,
                    0,
                    None,
                    70,
                    7,
                    None,
                    None,
                    None,
                    None,
                ],
            ),
            (
                ["ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "add", "add", "getProduct"],
                [[], [1], [1], [0], [5], [4], [1], [1], [0], [3]],
                [None, None, None, None, None, None, 4, None, None, 0],
            ),
            (
                ["ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct"],
                [[], [9], [4], [10], [5], [0], [1]],
                [None, None, None, None, None, None, 0],
            ),
            (
                [
                    "ProductOfNumbers",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "getProduct",
                    "add",
                    "getProduct",
                    "getProduct",
                    "add",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "getProduct",
                    "add",
                    "getProduct",
                ],
                [[], [2], [7], [10], [9], [3], [1], [5], [1], [1], [6], [10], [9], [2], [9], [2], [4], [1], [3]],
                [None, None, None, None, None, 630, 9, None, 5, 5, None, None, None, None, None, 18, 1620, None, 18],
            ),
            (
                ["ProductOfNumbers", "add", "add", "add", "add", "getProduct", "getProduct"],
                [[], [7], [8], [0], [1], [1], [2]],
                [None, None, None, None, None, 1, 0],
            ),
            (
                [
                    "ProductOfNumbers",
                    "add",
                    "add",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "getProduct",
                    "getProduct",
                    "add",
                    "getProduct",
                ],
                [[], [8], [2], [7], [7], [3], [7], [4], [4], [4], [0], [2]],
                [None, None, None, None, None, None, None, 1029, 1029, 1029, None, 0],
            ),
            (
                [
                    "ProductOfNumbers",
                    "add",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "add",
                    "getProduct",
                    "add",
                    "add",
                    "add",
                    "getProduct",
                    "add",
                    "getProduct",
                    "add",
                    "getProduct",
                    "add",
                    "getProduct",
                    "getProduct",
                    "getProduct",
                    "getProduct",
                ],
                [
                    [],
                    [3],
                    [7],
                    [3],
                    [6],
                    [2],
                    [0],
                    [3],
                    [3],
                    [4],
                    [7],
                    [3],
                    [3],
                    [3],
                    [8],
                    [3],
                    [8],
                    [4],
                    [2],
                    [4],
                    [4],
                ],
                [
                    None,
                    None,
                    None,
                    None,
                    None,
                    18,
                    None,
                    0,
                    None,
                    None,
                    None,
                    84,
                    None,
                    84,
                    None,
                    168,
                    None,
                    1344,
                    64,
                    1344,
                    1344,
                ],
            ),
            (
                ["ProductOfNumbers", "add", "add", "add", "add", "add"],
                [[], [0], [6], [7], [5], [5]],
                [None, None, None, None, None, None],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "ProductOfNumbers":
                        p = ProductOfNumbers()
                    case "add":
                        p.add(values[idx][0])
                    case "getProduct":
                        output = p.getProduct(values[idx][0])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
