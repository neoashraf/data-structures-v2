package rotation

func reverse(a *[]int, start, end int) {
	for start < end {
		(*a)[start], (*a)[end] = (*a)[end], (*a)[start]
		start++
		end--
	}
}

// ReversalAlgorithm - ...
func ReversalAlgorithm(a *[]int, d int) {
	length := len(*a)
	reverse(a, 0, d-1)
	reverse(a, d, length-1)
	reverse(a, 0, length-1)
}
