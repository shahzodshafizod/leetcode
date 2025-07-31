package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./pkg/ -run ^TestBST$
func TestBST(t *testing.T) {
	var bst BST

	for _, tc := range []struct {
		command string
		value   []int
		output  any
	}{
		{command: "BST", value: []int{}, output: nil},
		{command: "Search", value: []int{4}, output: false},
		{command: "Insert", value: []int{4}, output: nil},
		{command: "Search", value: []int{4}, output: true},
		{command: "Insert", value: []int{6}, output: nil},
		{command: "Search", value: []int{5}, output: false},
		{command: "Search", value: []int{6}, output: true},
		{command: "Search", value: []int{4}, output: true},
		{command: "Insert", value: []int{5}, output: nil},
		{command: "Search", value: []int{5}, output: true},
		{command: "Remove", value: []int{5}, output: nil},
		{command: "Search", value: []int{5}, output: false},
		{command: "Search", value: []int{4}, output: true},
		{command: "Search", value: []int{6}, output: true},
	} {
		var output any

		switch tc.command {
		case "BST":
			bst = NewBST()
		case "Search":
			output = bst.Search(tc.value[0])
		case "Insert":
			bst.Insert(tc.value[0])
		case "Remove":
			bst.Remove(tc.value[0])
		}

		assert.Equal(t, tc.output, output)
	}
}
