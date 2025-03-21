from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/

# python3 -m unittest graphs/2115-find-all-possible-recipes-from-given-supplies.py

class Solution(unittest.TestCase):
    # Approach: Topological Sort (Kahn's Algorithm)
    # Time: O(n+m+s), n=n=len(recipes), m=len(ingrediends), s=len(supplies)
    # Space: O(n+m+s)
    def findAllRecipes(self, recipes: List[str], ingredients: List[List[str]], supplies: List[str]) -> List[str]:
        indegree = defaultdict(int)
        deps = defaultdict(list)
        for idx, recipe in enumerate(recipes):
            indegree[recipe] += len(ingredients[idx])
            for ingredient in ingredients[idx]:
                deps[ingredient].append(recipe)
        result = []
        idx, n = 0, len(supplies)
        while idx < n:
            for recipe in deps[supplies[idx]]:
                indegree[recipe] -= 1
                if indegree[recipe] == 0:
                    result.append(recipe)
                    supplies.append(recipe)
                    n += 1
            idx += 1
        return result

    def test(self):
        for recipes, ingredients, supplies, expected in [
            (["bread"], [["yeast","flour"]], ["yeast","flour","corn"], ["bread"]),
            (["bread","sandwich"], [["yeast","flour"],["bread","meat"]], ["yeast","flour","meat"], ["bread","sandwich"]),
            # (["bread","sandwich","burger"], [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], ["yeast","flour","meat"], ["bread","sandwich","burger"]),
            # (["ju","fzjnm","x","e","zpmcz","h","q"], [["d"],["hveml","f","cpivl"],["cpivl","zpmcz","h","e","fzjnm","ju"],["cpivl","hveml","zpmcz","ju","h"],["h","fzjnm","e","q","x"],["d","hveml","cpivl","q","zpmcz","ju","e","x"],["f","hveml","cpivl"]], ["f","hveml","cpivl","d"], ["fzjnm","q","ju"]),
            (["bread","sandwich"], [["bread","flour"],["bread","flour"]], ["yeast","flour","meat"], []),
            # (["qxyj","vawos","nkov","bec","qiabz"], [["mxf"],["iy","qxyj","nkov","qiabz","bec"],["nw","xutnl","e"],["eep","km","nw","xutnl","e","iy","vawos","qxyj","qiabz"],["nyhyc"]], ["nw","eep","iy","e","xutnl","km"], ["nkov"]),
            (["bread","sandwich"], [["yeast","flour"],["bread","flour"]], ["yeast","flour","meat"], ["bread","sandwich"]),
        ]:
            output = self.findAllRecipes(recipes, ingredients, supplies)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
