package design

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct{}

func NewCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "N"
	}
	return fmt.Sprintf("%d,%s,%s",
		root.Val,
		c.serialize(root.Left),
		c.serialize(root.Right),
	)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	var vals = strings.Split(data, ",")
	var dfs func(idx int) (*TreeNode, int)
	dfs = func(idx int) (*TreeNode, int) {
		if vals[idx] == "N" {
			return nil, idx + 1
		}
		var node = &TreeNode{}
		node.Val, _ = strconv.Atoi(vals[idx])
		idx++
		node.Left, idx = dfs(idx)
		node.Right, idx = dfs(idx)
		return node, idx
	}
	var root, _ = dfs(0)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
