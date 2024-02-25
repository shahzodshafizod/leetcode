package design

// https://leetcode.com/problems/design-hashmap/

// 2. Chaining
type MyHashMap struct {
	array []*node[[2]int]
	cap   int
}

func NewMyHashMap() MyHashMap {
	const cap = 997
	var array = make([]*node[[2]int], cap)
	for i := 0; i < cap; i++ {
		array[i] = &node[[2]int]{} // dummy head node
	}
	return MyHashMap{array: array, cap: cap}
}

func (m *MyHashMap) Put(key int, value int) {
	curr := m.array[key%m.cap]
	for curr.next != nil {
		if curr.next.val[0] == key {
			curr.next.val[1] = value
			return
		}
		curr = curr.next
	}
	curr.next = &node[[2]int]{val: [2]int{key, value}}
}

func (m *MyHashMap) Get(key int) int {
	curr := m.array[key%m.cap].next
	for curr != nil {
		if curr.val[0] == key {
			return curr.val[1]
		}
		curr = curr.next
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	curr := m.array[key%m.cap]
	for curr.next != nil {
		if curr.next.val[0] == key {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

// // 1. Open Addressing (doesn't work)
// type MyHashMap struct {
// 	array []*[2]int
// 	len   int
// 	cap   int
// }

// func NewMyHashMap() MyHashMap {
// 	cap := 97 // >=7
// 	return MyHashMap{array: make([]*[2]int, cap), cap: cap}
// }

// func (m *MyHashMap) Put(key int, value int) {
// 	idx := key % m.cap
// 	for m.array[idx] != nil {
// 		if m.array[idx][0] == key {
// 			m.array[idx][1] = value
// 			return
// 		}
// 		idx++
// 		idx %= m.cap
// 	}
// 	m.array[idx] = &[2]int{key, value}
// 	m.len++
// 	if m.len >= m.cap {
// 		m.rehash()
// 	}
// }

// func (m *MyHashMap) Get(key int) int {
// 	if idx := m.getIndex(key); idx != -1 {
// 		return m.array[idx][1]
// 	}
// 	return -1
// }

// func (m *MyHashMap) Remove(key int) {
// 	if idx := m.getIndex(key); idx != -1 {
// 		m.array[idx] = nil
// 	}
// }

// func (m *MyHashMap) getIndex(key int) int {
// 	idx := key % m.cap
// 	for m.array[idx] != nil {
// 		if m.array[idx][0] == key {
// 			return idx
// 		}
// 		idx++
// 		idx %= m.cap
// 	}
// 	return -1
// }

// func (m *MyHashMap) rehash() {
// 	var array = m.array
// 	m.cap = m.nextPrime(m.cap * 3)
// 	m.array = make([]*[2]int, m.cap)
// 	m.len = 0
// 	for _, item := range array {
// 		if item != nil {
// 			m.Put(item[0], item[1])
// 		}
// 	}
// }

// func (m *MyHashMap) nextPrime(start int) int {
// 	for start%2 == 0 || start%3 == 0 || start%5 == 0 || start%7 == 0 {
// 		start++
// 	}
// 	return start
// }

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */