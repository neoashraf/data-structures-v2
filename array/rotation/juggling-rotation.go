package rotation

// GCD - ...
func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

// JugglingRotation - ...
func JugglingRotation(a *[]int, d int) {
	length := len(*a)
	juggle := GCD(length, d)

	for i := 0; i < juggle; i++ {
		temp := (*a)[i]
		j := i // j will change; i will not change

		for {
			k := j + d

			if k >= length {
				k -= length
			}

			if k == i {
				break
			}

			(*a)[j] = (*a)[k]
			j = k
		}

		(*a)[j] = temp
	}
}
