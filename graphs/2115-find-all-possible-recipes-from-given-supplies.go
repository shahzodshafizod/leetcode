package graphs

// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/

// Approach: Topological Sort (Kahn's Algorithm)
// Time: O(n+m+s), n=n=len(recipes), m=len(ingrediends), s=len(supplies)
// Space: O(n+m+s)
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	indegree := make(map[string]int)
	deps := make(map[string][]string)

	for idx, recipe := range recipes {
		indegree[recipe] = len(ingredients[idx])

		for _, ingredient := range ingredients[idx] {
			deps[ingredient] = append(deps[ingredient], recipe)
		}
	}

	result := make([]string, 0)

	for idx, n := 0, len(supplies); idx < n; idx++ {
		for _, recipe := range deps[supplies[idx]] {
			indegree[recipe]--
			if indegree[recipe] == 0 {
				result = append(result, recipe)
				supplies = append(supplies, recipe)
				n++
			}
		}
	}

	return result
}
