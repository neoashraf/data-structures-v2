package rotation

// RotateOneByOne - ...
func RotateOneByOne(a *[]int, d int) {
	length := len(*a)
	for i := 0; i < d; i++ {
		temp := (*a)[0]
		for j := 0; j < length-1; j++ {
			(*a)[j] = (*a)[j+1]
		}
		(*a)[length-1] = temp
	}
}
