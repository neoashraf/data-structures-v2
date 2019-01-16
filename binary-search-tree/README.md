# Binary Search Tree

* If a Binary tree is Binary Search Tree

``` go
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func isBST (root *Node, left *Node, right *Node) bool{
	if(root == nil){
		return true
	}
	if(left != nil && root.Value < left.Value){
		return false
	}
	if(Right != nil && root.Value > root.Value){
		return false
	}

	return isBST (root.Left, left, root) && isBST (root.Right, root, Right);
}

func main(){
	isBST(root, nil, nil);
}

```

* Searching

``` go
func searchRecursively(n *Node, searchTerm int) *Node{
	if n == nil  || n.Value == searchTerm { return n }
	if n.Value < searchTerm { return searchRecursively(n.Right, searchTerm) }
	return searchRecursively(n.Left, searchTerm)
}
```

``` go
func searchIteratively(n *Node, searchTerm int) *Node {
	currentNode := n
	while currentNode != nil {
		if searchTerm == currentNode.Value { return currentNode }
		else if searchTerm < currentNode.Value { currentNode = currentNode.Left } 
		currentNode = currentNode.Right
	}
	return currentNode
}
```

* Find Minimum value in BST

``` go
find searchMinimumValue(r *Node) interface{} {
	if r.Left == nil {return r.Value}
	return searchMinimumValue(r.Left)
}
```

* Find Maximum value in BST

``` go
find searchMaximumValue(r *Node) interface{} {
	if r.Right == nil {return r.Value}
	return searchMaximumValue(r.Right)
}
```
