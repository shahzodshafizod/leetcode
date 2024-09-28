package design

// https://leetcode.com/problems/design-hashset/

// 2. Chaining
type MyHashSet struct {
	array []*listNode[int]
	cap   int
}

func NewMyHashSet() MyHashSet {
	const cap = 997
	var array = make([]*listNode[int], cap)
	for idx := range array {
		array[idx] = &listNode[int]{} // dummy head node
	}
	return MyHashSet{array: array, cap: cap}
}

func (m *MyHashSet) Add(key int) {
	curr := m.array[key%m.cap]
	for curr.next != nil {
		if curr.next.val == key {
			return
		}
		curr = curr.next
	}
	curr.next = &listNode[int]{val: key}
}

func (m *MyHashSet) Remove(key int) {
	curr := m.array[key%m.cap]
	for curr.next != nil {
		if curr.next.val == key {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func (m *MyHashSet) Contains(key int) bool {
	curr := m.array[key%m.cap].next
	for curr != nil {
		if curr.val == key {
			return true
		}
		curr = curr.next
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

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
