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
* InOrder - Left, Root, Right

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

* PreOrder - Root, Left, Right

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

* PostOrder - Left, Right, Root

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
