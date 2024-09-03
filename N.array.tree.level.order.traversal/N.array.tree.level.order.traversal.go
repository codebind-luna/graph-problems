package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	vals := [][]int{}
	if root == nil {
		return vals
	}
	var queue [][]*Node
	queue = append(queue, []*Node{root})

	for len(queue) > 0 {
		elements := queue[0]
		queue = queue[1:]
		var val []int

		for _, ele := range elements {
			val = append(val, ele.Val)
		}

		if len(val) > 0 {
			vals = append(vals, val)
		}

		var list []*Node
		for _, ele := range elements {
			if len(ele.Children) > 0 {
				for _, child := range ele.Children {
					if child != nil {
						list = append(list, child)
					}
				}
			}

		}
		if len(list) > 0 {
			queue = append(queue, list)
		}
	}
	return vals
}
