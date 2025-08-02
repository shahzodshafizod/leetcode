package design

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/design-hashset/

// 2. Chaining
type MyHashSet struct {
	array []*pkg.ListNode
	cap   int
}

func NewMyHashSet() MyHashSet {
	const capacity = 997

	array := make([]*pkg.ListNode, capacity)
	for idx := range array {
		array[idx] = &pkg.ListNode{} // dummy head node
	}

	return MyHashSet{array: array, cap: capacity}
}

func (m *MyHashSet) Add(key int) {
	curr := m.array[key%m.cap]
	for curr.Next != nil {
		if curr.Next.Val == key {
			return
		}

		curr = curr.Next
	}

	curr.Next = &pkg.ListNode{Val: key}
}

func (m *MyHashSet) Remove(key int) {
	curr := m.array[key%m.cap]
	for curr.Next != nil {
		if curr.Next.Val == key {
			curr.Next = curr.Next.Next
			return
		}

		curr = curr.Next
	}
}

func (m *MyHashSet) Contains(key int) bool {
	curr := m.array[key%m.cap].Next
	for curr != nil {
		if curr.Val == key {
			return true
		}

		curr = curr.Next
	}

	return false
}

// // 1. Open Addressing (doesn't work)
// type MyHashSet struct {
// 	array []*int
// 	len   int
// 	cap   int
// }

// func NewMyHashSet() MyHashSet {
// 	const cap = 97 // >=7
// 	return MyHashSet{array: make([]*int, cap), cap: cap}
// }

// func (m *MyHashSet) Add(key int) {
// 	idx := key % m.cap
// 	for m.array[idx] != nil {
// 		if *m.array[idx] == key {
// 			return
// 		}
// 		idx++
// 		idx %= m.cap
// 	}
// 	m.array[idx] = &key
// 	m.len++
// 	if m.len >= m.cap {
// 		m.rehash()
// 	}
// }

// func (m *MyHashSet) Remove(key int) {
// 	if idx := m.getIndex(key); idx != -1 {
// 		m.array[idx] = nil
// 	}
// }

// func (m *MyHashSet) Contains(key int) bool {
// 	return m.getIndex(key) != -1
// }

// func (m *MyHashSet) getIndex(key int) int {
// 	idx := key % m.cap
// 	for m.array[idx] != nil {
// 		if *m.array[idx] == key {
// 			return idx
// 		}
// 		idx++
// 		idx %= m.cap
// 	}
// 	return -1
// }

// func (m *MyHashSet) rehash() {
// 	var array = m.array
// 	m.cap = m.nextPrime(m.cap * 3)
// 	m.array = make([]*int, m.cap)
// 	m.len = 0
// 	for _, key := range array {
// 		if key != nil {
// 			m.Add(*key)
// 		}
// 	}
// }

// func (m *MyHashSet) nextPrime(start int) int {
// 	for start%2 == 0 || start%3 == 0 || start%5 == 0 || start%7 == 0 {
// 		start++
// 	}
// 	return start
// }

/*
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
