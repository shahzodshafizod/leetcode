package design

/*
Tries

Problem:
Implement a trie (pronounced as "try") with insert, search, and startsWith methods.

interface Trie {
	void insert(String word);
	Boolean search(String word);
	Boolean startsWith(String prefix);
}

Step 1: Verify the constraints
	- Can we implement helper classes/objects?
		: Yes, you can use any features you see fit.
Step 2: Write out some test cases
	(e:[end])<--(l)<--(p:[end])<--(p)<--(a)<--(Root)-->(d)-->(o)-->(g:[end])
	                                                    |
	                                                    v
	                                                   (a)-->(s)-->(h:[end])
		insert("apple")
		search("dog") // false
		insert("dog")
		search("dog") // true
		startsWith("app") // true
		search("app") // false
		insert("app")
		search("app") // true
		insert("dash")
*/

// https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
	end      bool
	children map[byte]*Trie
}

func NewTrie() Trie {
	return Trie{end: false, children: make(map[byte]*Trie)}
}

func (t *Trie) Insert(word string) {
	t._insert(word, len(word), 0, t)
}

func (t *Trie) _insert(word string, n int, i int, node *Trie) {
	if i == n {
		node.end = true

		return
	}

	letter := word[i]

	next := node.children[letter]
	if next == nil {
		next = &Trie{end: false, children: make(map[byte]*Trie)}
		node.children[letter] = next
	}

	t._insert(word, n, i+1, next)
}

func (t *Trie) Search(word string) bool {
	return t._search(word, len(word), 0, t, false)
}

func (t *Trie) _search(word string, n int, i int, node *Trie, prefix bool) bool {
	if i == n {
		return prefix || node.end
	}

	if next := node.children[word[i]]; next != nil {
		return t._search(word, n, i+1, next, prefix)
	}

	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return t._search(prefix, len(prefix), 0, t, true)
}

// type Trie struct {
// 	end      bool
// 	children [26]*Trie
// }

// func NewTrie() Trie {
// 	return Trie{}
// }

// func (t *Trie) Insert(word string) {
// 	var current = t
// 	for _, r := range word {
// 		var idx = r - 'a'
// 		if current.children[idx] == nil {
// 			current.children[idx] = &Trie{}
// 		}
// 		current = current.children[idx]
// 	}
// 	current.end = true
// }

// func (t *Trie) Search(word string) bool {
// 	var current = t
// 	for _, r := range word {
// 		var idx = r - 'a'
// 		if current.children[idx] == nil {
// 			return false
// 		}
// 		current = current.children[idx]
// 	}
// 	return current.end
// }

// func (t *Trie) StartsWith(prefix string) bool {
// 	var current = t
// 	for _, r := range prefix {
// 		var idx = r - 'a'
// 		if current.children[idx] == nil {
// 			return false
// 		}
// 		current = current.children[idx]
// 	}
// 	return true
// }

/*
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
