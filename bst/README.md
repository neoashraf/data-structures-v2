# Binary Search Tree
* Searching

``` go
func searchRecursively(n *Node, searchTerm int) *Node{
	if n == nil  || n.Value == searchTerm { return n }
	if n.Value < searchTerm { return searchRecursively(n.Right, searchTerm) }
	return searchRecursively(n.Left, searchTerm)
}
```
