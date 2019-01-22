# Binary Tree

* Single Node Structure
``` go
type Node struct {
	Value interface{}
	Left *Node
	Right *Node
}
```
* Simple Binary Tree Construction
[Example Code](https://play.golang.org/p/yDOdO1wchyH)
``` go
package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	Head *Node
	Size int
}

func main() {
	n0 := Node{Value: 10}
	t := Tree{Head: &n0, Size: 1}

	n1 := Node{Value: 5}
	n2 := Node{Value: 25}
	n3 := Node{Value: 20}
	n4 := Node{Value: 6}

	t.Head.Left = &n1
	t.Head.Right = &n2

	t.Head.Right.Left = &n3
	t.Head.Left.Right = &n4
	inOrder(t.Head)
}

func inOrder(r *Node) {
	if r == nil {
		return
	}
	if r.Left != nil {
		inOrder(r.Left)
	}
	fmt.Println(r.Value)
	if r.Right != nil {
		inOrder(r.Right)
	}
}

/* 
Outputs 
------------
5 6 10 20 25
*/
```

## Traversal
* In Order - Left, Root, Right

``` go
func inOrder(r *Node) {
	if r == nil {
		return
	}
	if r.Left != nil {
		inOrder(r.Left)
	}
	fmt.Println(r.Value)
	if r.Right != nil {
		inOrder(r.Right)
	}
}
```

* Pre Order - Root, Left, Right

``` go
func preOrder(r *Node) {
	if r == nil {
		return
	}
	
	fmt.Println(r.Value)
	
	if r.Left != nil {
		inOrder(r.Left)
	}

	if r.Right != nil {
		inOrder(r.Right)
	}
}

```

* Post Order - Left, Right, Root

``` go
func postOrder(r *Node) {
	if r == nil {
		return
	}
	
	if r.Left != nil {
		inOrder(r.Left)
	}
	
	if r.Right != nil {
		inOrder(r.Right)
	}
	
	fmt.Println(r.Value)
}

```
* Size of a tree using Recursion
``` go
func sizeOfTree(r *Node) int{
	if r == nil {return 0}
	size := 1
	
	if r.Left != nil {
		size = size + sizeOfTree(r.Left)
	}
	
	if r.Right != nil {
		size = size + sizeOfTree(r.Right)
	}
	
	return size
}
```

* Size of a tree using Iteration

``` go
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
	
	for !q.IsEmpty(){
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

```
