package rotation

// SimpleRotation - ...
func SimpleRotation(a []int, n int) []int {
	length := len(a)
	b := make([]int, n)

	for p := 0; p < n; p++ {
		b[p] = a[p]
	}

	for i := 0; i < length-n; i++ {
		a[i] = a[(i + n)]
	}
	k := 0
	for j := length - n; j < length; j++ {
		a[j] = b[k]
		k++
	}
	return a
}
