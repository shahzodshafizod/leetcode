package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/design-task-manager/

type TaskManager struct {
	tasks      *pkg.Heap[[2]int] // [priority, taskID]
	priorities map[int][2]int    // taskID -> [priority, userID]
}

func NewTaskManager(tasks [][]int) TaskManager {
	t := TaskManager{
		tasks: pkg.NewHeap(
			make([][2]int, 0),
			func(x, y [2]int) bool {
				return x[0] > y[0] || x[0] == y[0] && x[1] > y[1]
			},
		),
		priorities: make(map[int][2]int),
	}

	var userID, taskID, priority int
	for _, task := range tasks {
		userID, taskID, priority = task[0], task[1], task[2]
		heap.Push(t.tasks, [2]int{priority, taskID})
		t.priorities[taskID] = [2]int{priority, userID}
	}

	return t
}

func (t *TaskManager) Add(userID int, taskID int, priority int) {
	heap.Push(t.tasks, [2]int{priority, taskID})
	t.priorities[taskID] = [2]int{priority, userID}
}

func (t *TaskManager) Edit(taskID int, newPriority int) {
	heap.Push(t.tasks, [2]int{newPriority, taskID})
	t.priorities[taskID] = [2]int{newPriority, t.priorities[taskID][1]}
}

func (t *TaskManager) Rmv(taskID int) {
	t.priorities[taskID] = [2]int{-1, t.priorities[taskID][1]}
}

func (t *TaskManager) ExecTop() int {
	for t.tasks.Len() > 0 && t.tasks.Peak()[0] != t.priorities[t.tasks.Peak()[1]][0] {
		heap.Pop(t.tasks)
	}

	userID := -1

	if t.tasks.Len() > 0 {
		item, ok := heap.Pop(t.tasks).([2]int)
		_ = ok
		userID = t.priorities[item[1]][1]
		t.Rmv(item[1])
	}

	return userID
}

// /**
//  * Your TaskManager object will be instantiated and called as such:
//  * obj := Constructor(tasks);
//  * obj.Add(userId,taskId,priority);
//  * obj.Edit(taskId,newPriority);
//  * obj.Rmv(taskId);
//  * param_4 := obj.ExecTop();
//  */
