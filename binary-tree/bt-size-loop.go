package main

import (
	"fmt"
)

type QNode struct {
	i *Node
	n *QNode
}

type Queue []*QNode

func (q *Queue) Enqueue(p *Node) {
	qn := &QNode{i: p}
	*q = append(*q, qn)
}
func (q *Queue) Dequeue() *Node {
	if len(*q) == 0 {
		return nil
	}
	qn := (*q)[0]
	*q = (*q)[1:len(*q)]
	return qn.i
}

func (q *Queue) Front() *Node {
	if len(*q) == 0 {
		return nil
	}
	qn := (*q)[0]
	return qn.i
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BTree struct {
	Head *Node
}

func (b *BTree) sizeOfTree() int {
	r := b.Head
	i := 0
	q := Queue{}
	q.Enqueue(r)
	// fmt.Println((q.Front()).Value)

	for !q.IsEmpty() {
		i += 1
		c := q.Dequeue()
		fmt.Println(c.Value)
		if c.Left != nil {
			q.Enqueue(c.Left)
		}

		if c.Right != nil {
			q.Enqueue(c.Right)
		}
	}
	return i
}

func main() {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	n5 := &Node{Value: 5}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n3.Right = n5

	b := &BTree{n1}
	fmt.Println("Size -> ", b.sizeOfTree())
}
