from collections import defaultdict
from typing import List, Any, Optional, Dict, Tuple
import unittest
from sortedcontainers import SortedList

# https://leetcode.com/problems/design-movie-rental-system/

# python3 -m unittest design/1912-design-movie-rental-system.py


class MovieRentingSystem:

    def __init__(self, n: int, entries: List[List[int]]):
        _ = n
        self.shops: Dict[int, SortedList[Tuple[int, int]]] = defaultdict(SortedList)  # movie -> [(price, shop)]
        self.rented: SortedList[Tuple[int, int, int]] = SortedList()  # (price, shop, movie)
        self.prices: Dict[Tuple[int, int], int] = {}  # (shop, movie) -> price
        for shop, movie, price in entries:
            self.shops[movie].add((price, shop))
            self.prices[(shop, movie)] = price

    def search(self, movie: int) -> List[int]:
        shops: List[int] = []
        for _, shop in self.shops[movie]:
            shops.append(shop)
            if len(shops) == 5:
                break
        return shops

    def rent(self, shop: int, movie: int) -> None:
        price = self.prices[(shop, movie)]
        self.shops[movie].discard((price, shop))
        self.rented.add((price, shop, movie))

    def drop(self, shop: int, movie: int) -> None:
        price = self.prices[(shop, movie)]
        self.rented.discard((price, shop, movie))
        self.shops[movie].add((price, shop))

    def report(self) -> List[List[int]]:
        movies: List[List[int]] = []
        for _, shop, movie in self.rented:
            movies.append([shop, movie])
            if len(movies) == 5:
                break
        return movies


# Your MovieRentingSystem object will be instantiated and called as such:
# obj = MovieRentingSystem(n, entries)
# param_1 = obj.search(movie)
# obj.rent(shop,movie)
# obj.drop(shop,movie)
# param_4 = obj.report()


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[tuple[List[str], List[Any], List[Any]]] = [
            (
                ["MovieRentingSystem", "search", "rent", "rent", "report", "drop", "search"],
                [[3, [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]]], [1], [0, 1], [1, 2], [], [1, 2], [2]],
                [None, [1, 0, 2], None, None, [[0, 1], [1, 2]], None, [0, 1]],
            ),
        ]
        for commands, values, expected in test_cases:
            renting: Optional[MovieRentingSystem] = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "MovieRentingSystem":
                        renting = MovieRentingSystem(values[idx][0], values[idx][1])
                    case "search":
                        if renting is not None:
                            output = renting.search(values[idx][0])
                    case "rent":
                        if renting is not None:
                            renting.rent(values[idx][0], values[idx][1])
                    case "drop":
                        if renting is not None:
                            renting.drop(values[idx][0], values[idx][1])
                    case "report":
                        if renting is not None:
                            output = renting.report()
                    case _:
                        pass
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
