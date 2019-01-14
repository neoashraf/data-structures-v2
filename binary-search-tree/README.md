# Binary Search Tree
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
