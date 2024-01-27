package graphs

/*
Problem:
A company has n employees with unique IDs from 0 to n-1. The head of the company has the ID headID.

You will  receive a managers array where managers[i] is the ID of the manager for employee i.
Each employee has one direct manager. The company head has no manager (managers[headID] = -1).
It's guaranteed the subordination relationships will have a tree structure.

example: 8 employees: 0, 1, 2, 3, 4, 5, 6, 7, headID: 4, managers: [2, 2, 4, 6, -1, 4, 4, 5]
			 +----------(4)----------+
			 |			 |			 |
		 +--(5)		 +--(2)--+		(6)--+
		 |			 |		 |			 |
		(7)			(0)		(1)			(3)

The head of the company wants to inform all employees of news. He will inform his direct subordinates
who will inform their direct subordinates and so on until everyone knows the news.

You will receive an informTime array where informTime[i] is the time
it takes for employee i to inform all their direct subordinates.
Return the total number of minutes it takes to inform all employees of the news.

example: informTime = [0, 0, 4, 0, 7, 3, 6, 0]
			 +----------(7)----------+
			 |			 |			 |
		 +--(3)		 +--(4)--+		(6)--+
		 |			 |		 |			 |
		(0)			(0)		(0)			(0)

numOfMinutes: 13 (max(10, 11, 13))

Step 1: Verify the constraints
	- Is it a cyclic graph?
		: No.
	- Unconnected?
		: No.
	- Weighted?
	- Directed?
	- Can employees have more than 1 manager?
		: No, employees can only have 1 manager.
	- Does every employee have a manager?
		: Yes, every employee has a manager except the head of the company, who has no manager.
Step 2: Write out some test cases
	- n=1, headID=0, managers=[-1], informTime=[0], minutes=0
	- n=7, headID=6, managers=[1,2,3,4,5,6,-1], informTime=[0,6,5,4,3,2,1], minutes=21
	- n=8, headID=4, managers=[2,2,4,6,-1,4,4,5], informTime=[0,0,4,0,7,3,6,0], minutes=13
*/

// https://leetcode.com/problems/time-needed-to-inform-all-employees/

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var adjacencyList = make([][]int, n)
	for id := 0; id < n; id++ {
		if manager[id] != -1 {
			if adjacencyList[manager[id]] == nil {
				adjacencyList[manager[id]] = []int{id}
			} else {
				adjacencyList[manager[id]] = append(adjacencyList[manager[id]], id)
			}
		}
	}
	return numOfMinutesDFS(adjacencyList, informTime, headID)
}

func numOfMinutesDFS(adjacencyList [][]int, informTime []int, ID int) int {
	if adjacencyList[ID] == nil {
		return 0
	}
	var maximum = 0
	for _, employeeID := range adjacencyList[ID] {
		maximum = max(maximum, numOfMinutesDFS(adjacencyList, informTime, employeeID))
	}
	return maximum + informTime[ID]
}
