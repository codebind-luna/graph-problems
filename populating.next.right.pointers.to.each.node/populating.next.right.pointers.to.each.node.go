package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type queue []*Node

func (q *queue) push(n *Node) {
	(*q) = append((*q), n)
}

func (q *queue) pop() *Node {
	t := (*q)[0]
	(*q) = (*q)[1:]
	return t
}

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := new(queue)
	q.push(root)
	for !q.isEmpty() {

		var prev *Node
		q1 := new(queue)

		for !q.isEmpty() {
			t := q.pop()
			t.Next = prev
			prev = t

			if t.Right != nil {
				q1.push(t.Right)
			}

			if t.Left != nil {
				q1.push(t.Left)
			}
		}

		for !q1.isEmpty() {
			q.push(q1.pop())
		}
	}

	return root
}
