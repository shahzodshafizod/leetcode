package dp

// https://leetcode.com/problems/construct-string-with-minimum-cost/

// Approach 5: Aho-Corasick Automaton
// Time: O(N + M) where N = len(target), M = sum(len(word) for word in words)
// Space: O(M) for the trie + O(N) for dp array
func minimumCost(target string, words []string, costs []int) int {
	// TrieNode represents a node in the Aho-Corasick automaton
	type TrieNode struct {
		children map[rune]*TrieNode
		fail     *TrieNode // Failure link
		outputs  [][2]int  // List of [word_length, min_cost] for words ending here
	}

	newTrieNode := func() *TrieNode {
		return &TrieNode{
			children: make(map[rune]*TrieNode),
			outputs:  make([][2]int, 0),
		}
	}

	// Step 1: Build trie with minimum costs for duplicate words
	wordCosts := make(map[string]int)

	// Track minimum cost for each unique word
	for i, word := range words {
		if pcost, exists := wordCosts[word]; !exists || costs[i] < pcost {
			wordCosts[word] = costs[i]
		}
	}

	// Insert words into trie
	root := newTrieNode()
	for word, cost := range wordCosts {
		curr := root
		for _, c := range word {
			if _, exists := curr.children[c]; !exists {
				curr.children[c] = newTrieNode()
			}

			curr = curr.children[c]
		}

		curr.outputs = append(curr.outputs, [2]int{len(word), cost})
	}

	// Step 2: Build failure links (BFS)
	queue := make([]*TrieNode, 0)

	// Initialize first level failure links to root
	for _, child := range root.children {
		child.fail = root
		queue = append(queue, child)
	}

	// BFS to build failure links
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for c, child := range curr.children {
			queue = append(queue, child)

			// Find failure link
			failNode := curr.fail
			for failNode != nil {
				if _, exists := failNode.children[c]; exists {
					child.fail = failNode.children[c]

					break
				}

				failNode = failNode.fail
			}

			if failNode == nil {
				child.fail = root
			}

			// Inherit outputs from failure link
			child.outputs = append(child.outputs, child.fail.outputs...)
		}
	}

	// Step 3: DP with Aho-Corasick matching
	const inf = int(1e9)

	n := len(target)

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = inf
	}

	dp[0] = 0

	curr := root

	for i := range n {
		c := rune(target[i])

		// Follow failure links until we find a match or reach root
		// This ensures we never backtrack in the target string - we only move
		// forward through target while using failure links to navigate the trie.
		for curr != nil {
			if _, exists := curr.children[c]; exists {
				curr = curr.children[c]

				break
			}

			curr = curr.fail
		}

		if curr == nil {
			curr = root
		}

		// Process all matching words ending at position i+1
		for _, output := range curr.outputs {
			wordLen, cost := output[0], output[1]

			start := i - wordLen + 1
			if dp[start] != inf {
				dp[i+1] = min(dp[i+1], dp[start]+cost)
			}
		}
	}

	if dp[n] < inf {
		return dp[n]
	}

	return -1
}

// // Approach 4: Bottom-up DP + Trie (Time Limit Exceeded)
// // Time: O(n * L_max) where n is target length, L_max is longest word
// // Space: O(n + total_chars) for DP array and Trie
// func minimumCost(target string, words []string, costs []int) int {
// 	type node struct {
// 		cost     int
// 		children map[string]*node
// 	}

// 	newNode := func() *node {
// 		return &node{
// 			cost:     -1,
// 			children: make(map[string]*node),
// 		}
// 	}
// 	root := newNode()
// 	add := func(word string, cost int) {
// 		curr := root

// 		for _, c := range word {
// 			key := string(c)
// 			if curr.children[key] == nil {
// 				curr.children[key] = newNode()
// 			}

// 			curr = curr.children[key]
// 		}

// 		if curr.cost == -1 || cost < curr.cost {
// 			curr.cost = cost
// 		}
// 	}

// 	for i := range words {
// 		add(words[i], costs[i])
// 	}

// 	n := len(target)

// 	const maximum = (1 << 32) - 1

// 	dp := make([]int, n+1)
// 	for i := range n + 1 {
// 		dp[i] = maximum
// 	}

// 	dp[0] = 0

// 	for i := range n {
// 		if dp[i] == maximum {
// 			continue // Can't reach this position
// 		}

// 		curr := root

// 		for j := i; j < n; j++ {
// 			key := string(target[j])
// 			if curr.children[key] == nil {
// 				break
// 			}

// 			curr = curr.children[key]
// 			if curr.cost != -1 {
// 				dp[j+1] = min(dp[j+1], dp[i]+curr.cost)
// 			}
// 		}
// 	}

// 	if dp[n] == maximum {
// 		return -1
// 	}

// 	return dp[n]
// }

// // Approach 3: Top-Down DP + Trie (Time Limit Exceeded)
// // Time: O(n * L_max) where n is target length, L_max is longest word
// // Space: O(n + total_chars) for DP array and Trie
// func minimumCost(target string, words []string, costs []int) int {
// 	type node struct {
// 		cost     int
// 		children map[string]*node
// 	}

// 	newNode := func() *node {
// 		return &node{
// 			cost:     -1,
// 			children: make(map[string]*node),
// 		}
// 	}
// 	root := newNode()
// 	add := func(word string, cost int) {
// 		curr := root

// 		for _, c := range word {
// 			key := string(c)
// 			if curr.children[key] == nil {
// 				curr.children[key] = newNode()
// 			}

// 			curr = curr.children[key]
// 		}

// 		if curr.cost == -1 || cost < curr.cost {
// 			curr.cost = cost
// 		}
// 	}

// 	for i := range words {
// 		add(words[i], costs[i])
// 	}

// 	n := len(target)
// 	memo := make([]*int, n)

// 	var dp func(pos int) int

// 	dp = func(pos int) int {
// 		if pos == n {
// 			return 0
// 		}

// 		if memo[pos] != nil {
// 			return *memo[pos]
// 		}

// 		minCost := -1

// 		curr := root

// 		for j := pos; j < n; j++ {
// 			key := string(target[j])
// 			if curr.children[key] == nil {
// 				break
// 			}

// 			curr = curr.children[key]
// 			if curr.cost != -1 {
// 				cost := dp(j + 1)
// 				if cost != -1 && (minCost == -1 || curr.cost+cost < minCost) {
// 					minCost = curr.cost + cost
// 				}
// 			}
// 		}

// 		memo[pos] = &minCost

// 		return minCost
// 	}

// 	return dp(0)
// }

// // Approach 2: Bottom-up DP (Tabulation)
// // Time: O(n * m * L) where n is target length, m is words count, L is avg word length
// // Space: O(n) for DP array
// func minimumCost(target string, words []string, costs []int) int {
// 	n := len(target)
// 	// dp[i] = minimum cost to construct target[0:i]
// 	dp := make([]int, n+1)
// 	for i := range n {
// 		dp[i+1] = math.MaxInt32
// 	}

// 	dp[0] = 0 // Empty string costs 0

// 	// For each position in target
// 	for i := 0; i <= n; i++ {
// 		if dp[i] == math.MaxInt32 {
// 			continue // Can't reach this position
// 		}

// 		// Try each word
// 		for j, word := range words {
// 			wordLen := len(word)
// 			// Check if word can fit and matches target at position i
// 			if i+wordLen <= n && target[i:i+wordLen] == word {
// 				// Update minimum cost to reach position i + wordLen
// 				dp[i+wordLen] = min(dp[i+wordLen], dp[i]+costs[j])
// 			}
// 		}
// 	}

// 	if dp[n] == math.MaxInt32 {
// 		return -1
// 	}

// 	return dp[n]
// }

// // Approach 1: Brute-force with Backtracking
// // Time: O(2^n * m) where n is target length, m is number of words
// // Space: O(n) for recursion stack
// func minimumCost(target string, words []string, costs []int) int {
// 	n := len(target)

// 	var backtrack func(pos int) int

// 	backtrack = func(pos int) int {
// 		if pos == n {
// 			return 0
// 		}

// 		minCost := -1

// 		var wlen, cost int
// 		for i, word := range words {
// 			wlen = len(word)
// 			if pos+wlen <= n && target[pos:pos+wlen] == word {
// 				cost = backtrack(pos + wlen)
// 				if cost != -1 && (minCost == -1 || costs[i]+cost < minCost) {
// 					minCost = costs[i] + cost
// 				}
// 			}
// 		}

// 		return minCost
// 	}

// 	return backtrack(0)
// }
