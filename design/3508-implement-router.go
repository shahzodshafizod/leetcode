package design

import (
	"container/list"
)

// https://leetcode.com/problems/implement-router/

type Router struct {
	deque  *list.List
	cap    int
	offset map[int]int
	unique map[[3]int]bool
	dests  map[int][]int
}

func NewRouter(memoryLimit int) Router {
	return Router{
		deque:  list.New(),
		cap:    memoryLimit,
		offset: make(map[int]int),
		unique: make(map[[3]int]bool),
		dests:  make(map[int][]int),
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	packet := [3]int{source, destination, timestamp}

	if r.unique[packet] {
		return false
	}

	if r.deque.Len() == r.cap {
		_ = r.ForwardPacket()
	}

	r.deque.PushBack(packet)
	r.unique[packet] = true
	r.dests[destination] = append(r.dests[destination], timestamp)

	return true
}

func (r *Router) ForwardPacket() []int {
	if r.deque.Len() == 0 {
		return []int{}
	}

	packet, ok := r.deque.Remove(r.deque.Front()).([3]int)
	_ = ok

	r.offset[packet[1]]++
	r.unique[packet] = false

	return []int{packet[0], packet[1], packet[2]}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	left, right := r.offset[destination], len(r.dests[destination])-1

	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if r.dests[destination][mid] < startTime {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	startIndex := left

	left, right = r.offset[destination], len(r.dests[destination])-1

	for left <= right {
		mid = left + (right-left)/2
		if r.dests[destination][mid] > endTime {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	endIndex := right

	return endIndex - startIndex + 1
}

// /**
//  * Your Router object will be instantiated and called as such:
//  * obj := Constructor(memoryLimit);
//  * param_1 := obj.AddPacket(source,destination,timestamp);
//  * param_2 := obj.ForwardPacket();
//  * param_3 := obj.GetCount(destination,startTime,endTime);
//  */
