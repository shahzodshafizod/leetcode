package stacks

import (
	"sort"
	"strconv"
	"unicode"
)

// https://leetcode.com/problems/number-of-atoms/

// Approach: Parse Backward: Stack + Array + Sorting
// time: O(N Log N)
// space: O(N)
func countOfAtoms(formula string) string {
	type element struct {
		name  string
		atoms int
	}
	elem := &element{name: "", atoms: 0}
	dec := 1
	var r rune
	elements := make([]*element, 0)
	stack := make([]int, 0)
	multiply := 1
	for i := len(formula) - 1; i >= 0; i-- { // O(N)
		r = rune(formula[i])
		switch {
		case unicode.IsDigit(r):
			elem.atoms += int(r-'0') * dec
			dec *= 10
		case unicode.IsLower(r):
			elem.name = string(r) + elem.name
		case unicode.IsUpper(r):
			elem.name = string(r) + elem.name
			if elem.atoms == 0 {
				elem.atoms = 1
			}
			elem.atoms *= multiply
			elements = append(elements, elem)
			elem = &element{name: "", atoms: 0}
			dec = 1
		case r == ')':
			if elem.atoms == 0 {
				elem.atoms = 1
			}
			stack = append(stack, elem.atoms)
			multiply *= elem.atoms
			elem.atoms = 0
			dec = 1
		case r == '(':
			multiply /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	sort.Slice(elements, // O(N Log N)
		func(i, j int) bool {
			return elements[i].name < elements[j].name
		},
	)
	result := ""
	n := len(elements)
	for idx := range elements { // O(N)
		if idx+1 < n && elements[idx].name == elements[idx+1].name {
			elements[idx+1].atoms += elements[idx].atoms
		} else {
			result += elements[idx].name
			if elements[idx].atoms > 1 {
				result += strconv.Itoa(elements[idx].atoms)
			}
		}
	}
	return result
}

// // Approach: Stack + Hash + Array + Sorting
// // time: O(N Log N)
// // space: O(N)
// func countOfAtoms(formula string) string {
// 	type element struct {
// 		name  string
// 		atoms int
// 	}
// 	type node struct {
// 		val  *element
// 		next *node
// 	}
// 	// 1. parse into stack (stack is cosy for getting access to the last element)
// 	var stack *node = nil
// 	for _, r := range formula { // O(N)
// 		switch {
// 		case unicode.IsUpper(r) || r == '(' || r == ')':
// 			stack = &node{&element{string(r), 0}, stack}
// 		case unicode.IsLower(r):
// 			stack.val.name += string(r)
// 		case unicode.IsDigit(r):
// 			stack.val.atoms = stack.val.atoms*10 + int(r-'0')
// 		}
// 	}
// 	// 2. union all occurences and their atoms
// 	var set = make(map[string]int)
// 	var multiply = 1
// 	var mStack = make([]int, 0)
// 	for stack != nil { // O(N)
// 		switch stack.val.name {
// 		case ")":
// 			multiply *= stack.val.atoms
// 			mStack = append(mStack, stack.val.atoms)
// 		case "(":
// 			stack.val.atoms = mStack[len(mStack)-1]
// 			mStack = mStack[:len(mStack)-1]
// 			multiply /= stack.val.atoms
// 		default:
// 			if stack.val.atoms == 0 {
// 				stack.val.atoms = 1
// 			}
// 			set[stack.val.name] += stack.val.atoms * multiply
// 		}
// 		stack = stack.next
// 	}
// 	// 3. convert set to slice to sort the elements
// 	var elements = make([]*element, 0, len(set))
// 	for name, atoms := range set {
// 		elements = append(elements, &element{name, atoms})
// 	}
// 	sort.Slice(elements, // O(N Log N)
// 		func(i, j int) bool {
// 			return elements[i].name < elements[j].name
// 		},
// 	)
// 	// 4. build the result
// 	var result = ""
// 	for _, element := range elements { // O(N)
// 		result += element.name
// 		if element.atoms > 1 {
// 			result += strconv.Itoa(element.atoms)
// 		}
// 	}
// 	return result
// }
