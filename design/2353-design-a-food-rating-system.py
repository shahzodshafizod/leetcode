from collections import defaultdict
from typing import List, Dict, Any, Tuple, Optional
import unittest
import heapq

# from sortedcontainers import SortedSet

# https://leetcode.com/problems/design-a-food-rating-system/

# python3 -m unittest design/2353-design-a-food-rating-system.py


# # Approach: SortedSet
# class FoodRatings:

#     def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
#         self.cuisines: Dict[str, SortedSet[Tuple[int, str]]] = defaultdict(SortedSet)
#         self.foods: Dict[str, Tuple[int, str]] = {}
#         for food, cuisine, rating in zip(foods, cuisines, ratings):
#             self.cuisines[cuisine].add((-rating, food))
#             self.foods[food] = (rating, cuisine)

#     def changeRating(self, food: str, newRating: int) -> None:
#         rating, cuisine = self.foods[food]
#         self.cuisines[cuisine].remove((-rating, food))
#         self.cuisines[cuisine].add((-newRating, food))
#         self.foods[food] = (newRating, cuisine)

#     def highestRated(self, cuisine: str) -> str:
#         return self.cuisines[cuisine][0][1]


# Approach: Heap (Priority Queue)
class FoodRatings:

    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.cuisines: Dict[str, List[Tuple[int, str]]] = defaultdict(list)
        self.foods: Dict[str, Tuple[int, str]] = {}
        for food, cuisine, rating in zip(foods, cuisines, ratings):
            heapq.heappush(self.cuisines[cuisine], (-rating, food))
            self.foods[food] = (rating, cuisine)

    def changeRating(self, food: str, newRating: int) -> None:
        cuisine = self.foods[food][1]
        heapq.heappush(self.cuisines[cuisine], (-newRating, food))
        self.foods[food] = (newRating, cuisine)

    def highestRated(self, cuisine: str) -> str:
        while self.cuisines[cuisine][0][0] != -self.foods[self.cuisines[cuisine][0][1]][0]:
            heapq.heappop(self.cuisines[cuisine])
        return self.cuisines[cuisine][0][1]


# Your FoodRatings object will be instantiated and called as such:
# obj = FoodRatings(foods, cuisines, ratings)
# obj.changeRating(food,newRating)
# param_2 = obj.highestRated(cuisine)


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[tuple[List[str], List[Any], List[Any]]] = [
            (
                ["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"],
                [
                    [["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]],
                    ["korean"],
                    ["japanese"],
                    ["sushi", 16],
                    ["japanese"],
                    ["ramen", 16],
                    ["japanese"],
                ],
                [None, "kimchi", "ramen", None, "sushi", None, "ramen"],
            ),
        ]
        for commands, values, expected in test_cases:
            food_ratings: Optional[FoodRatings] = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "FoodRatings":
                        foods: List[str] = values[idx][0]
                        cuisines: List[str] = values[idx][1]
                        ratings: List[int] = values[idx][2]
                        food_ratings = FoodRatings(foods, cuisines, ratings)
                    case "changeRating":
                        if food_ratings is not None:
                            food_ratings.changeRating(values[idx][0], values[idx][1])
                    case "highestRated":
                        if food_ratings is not None:
                            output = food_ratings.highestRated(values[idx][0])
                    case _:
                        pass
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
