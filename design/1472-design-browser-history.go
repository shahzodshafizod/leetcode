package design

// https://leetcode.com/problems/design-browser-history/

// Approach#4: Array

type BrowserHistory struct {
	history []string
	current int
	size    int
}

func NewBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{history: []string{homepage}, current: 0, size: 1}
}

func (b *BrowserHistory) Visit(url string) {
	b.history = b.history[:b.current+1]
	b.history = append(b.history, url)
	b.current++
	b.size = b.current + 1
}

func (b *BrowserHistory) Back(steps int) string {
	if steps >= b.current {
		b.current = 0
	} else {
		b.current -= steps
	}
	return b.history[b.current]
}

func (b *BrowserHistory) Forward(steps int) string {
	if steps >= b.size-b.current {
		b.current = b.size - 1
	} else {
		b.current += steps
	}
	return b.history[b.current]
}

// // Approach#3: Two Stacks

// type BrowserHistory struct {
// 	before Stack[string]
// 	after  Stack[string]
// }

// func NewBrowserHistory(homepage string) BrowserHistory {
// 	before := NewStack[string]()
// 	before.Push(homepage)
// 	return BrowserHistory{before: before, after: NewStack[string]()}
// }

// func (b *BrowserHistory) Visit(url string) {
// 	b.after = NewStack[string]()
// 	b.before.Push(url)
// }

// func (b *BrowserHistory) Back(steps int) string {
// 	for steps > 0 && b.before.Size() > 1 {
// 		steps--
// 		b.after.Push(b.before.Pop())
// 	}
// 	return b.before.Peek()
// }

// func (b *BrowserHistory) Forward(steps int) string {
// 	for steps > 0 && !b.after.Empty() {
// 		steps--
// 		b.before.Push(b.after.Pop())
// 	}
// 	return b.before.Peek()
// }

// // Approach#2: Linked List

// type BrowserHistory struct {
// 	history []*node[string]
// 	current int
// 	size    int
// }

// func NewBrowserHistory(homepage string) BrowserHistory {
// 	return BrowserHistory{history: []*node[string]{{val: homepage}}, current: 0, size: 1}
// }

// func (b *BrowserHistory) Visit(url string) {
// 	b.history[b.current].next = &node[string]{val: url}
// 	b.history = b.history[:b.current+1]
// 	b.history = append(b.history, b.history[b.current].next)
// 	b.current++
// 	b.size = b.current + 1
// }

// func (b *BrowserHistory) Back(steps int) string {
// 	if steps >= b.current {
// 		b.current = 0
// 	} else {
// 		b.current -= steps
// 	}
// 	return b.history[b.current].val
// }

// func (b *BrowserHistory) Forward(steps int) string {
// 	if steps >= b.size-b.current {
// 		b.current = b.size - 1
// 	} else {
// 		b.current += steps
// 	}
// 	return b.history[b.current].val
// }

// func (b *BrowserHistory) Println(command string, values []any) {
// 	fmt.Printf("\nBROWSER HISTORY: command=[%s], values=%v\n", command, values)
// 	for head := b.history[0]; head != nil; head = head.next {
// 		if head == b.history[b.current] {
// 			fmt.Printf("[\"%s\"] -> ", head.val)
// 		} else {
// 			fmt.Printf("\"%s\" -> ", head.val)
// 		}
// 	}
// 	fmt.Println("nil")
// 	for current := 0; current < b.size; current++ {
// 		if current == b.current {
// 			fmt.Printf("[\"%s\"] -> ", b.history[current].val)
// 		} else {
// 			fmt.Printf("\"%s\" -> ", b.history[current].val)
// 		}
// 	}
// 	fmt.Println("nil")
// }

// // Approach#1: Doubly-Linked List

// type BrowserHistory struct {
// 	current *listNode[string]
// }

// func NewBrowserHistory(homepage string) BrowserHistory {
// 	return BrowserHistory{&listNode[string]{val: homepage}}
// }

// func (b *BrowserHistory) Visit(url string) {
// 	var newPage = &listNode[string]{val: url, prev: b.current}
// 	if b.current == nil {
// 		b.current = newPage
// 	} else {
// 		b.current.next = newPage
// 		b.current = newPage
// 	}
// }

// func (b *BrowserHistory) Back(steps int) string {
// 	if b.current == nil {
// 		return ""
// 	}
// 	for steps > 0 && b.current != nil && b.current.prev != nil {
// 		steps--
// 		b.current = b.current.prev
// 	}
// 	return b.current.val
// }

// func (b *BrowserHistory) Forward(steps int) string {
// 	if b.current == nil {
// 		return ""
// 	}
// 	for steps > 0 && b.current != nil && b.current.next != nil {
// 		steps--
// 		b.current = b.current.next
// 	}
// 	return b.current.val
// }

// func (b *BrowserHistory) Println(command string, values []any) {
// 	fmt.Printf("\nBROWSER HISTORY: command=[%s], values=%v\n", command, values)
// 	var head = b.current
// 	for head != nil && head.prev != nil {
// 		head = head.prev
// 	}
// 	for ; head != nil; head = head.next {
// 		if head == b.current {
// 			fmt.Printf("[\"%s\"] -> ", head.val)
// 		} else {
// 			fmt.Printf("\"%s\" -> ", head.val)
// 		}
// 	}
// 	fmt.Println("nil")
// }

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
