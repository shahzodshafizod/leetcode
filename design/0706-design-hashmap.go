package design

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/design-hashmap/

// 2. Chaining
type MyHashMap struct {
	keys []*pkg.ListNode
	vals []*pkg.ListNode
	cap  int
}

func NewMyHashMap() MyHashMap {
	const capacity = 997

	keys := make([]*pkg.ListNode, capacity)
	vals := make([]*pkg.ListNode, capacity)

	for i := 0; i < capacity; i++ {
		keys[i] = &pkg.ListNode{} // dummy head node
		vals[i] = &pkg.ListNode{} // dummy head node
	}

	return MyHashMap{keys: keys, vals: vals, cap: capacity}
}

func (m *MyHashMap) Put(key int, value int) {
	currKey := m.keys[key%m.cap]
	currVal := m.vals[key%m.cap]

	for currKey.Next != nil {
		if currKey.Next.Val == key {
			currVal.Next.Val = value

			return
		}

		currKey = currKey.Next
		currVal = currVal.Next
	}

	currKey.Next = &pkg.ListNode{Val: key}
	currVal.Next = &pkg.ListNode{Val: value}
}

func (m *MyHashMap) Get(key int) int {
	currKey := m.keys[key%m.cap].Next
	currVal := m.vals[key%m.cap].Next

	for currKey != nil {
		if currKey.Val == key {
			return currVal.Val
		}

		currKey = currKey.Next
		currVal = currVal.Next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	currKey := m.keys[key%m.cap]
	currVal := m.vals[key%m.cap]

	for currKey.Next != nil {
		if currKey.Next.Val == key {
			currKey.Next = currKey.Next.Next
			currVal.Next = currVal.Next.Next

			return
		}

		currKey = currKey.Next
		currVal = currVal.Next
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

/*
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
