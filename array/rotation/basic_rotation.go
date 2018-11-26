package rotation

import "fmt"

func rotateMe(arr *[]int, d int) {
	length := len(*arr)

	var next int

	for i := 0; i < length; i++ {

		copy := make([]int, d)

		for j := d; j < d*2; j++ {
			copy = append(copy, (*arr)[j])
		}

		for j := i; j < d+i; j++ {
			index := (j + d) % length
			(*arr)[index] = (*arr)[j]
		}

		for j := d*2; j < d*3; j++ {
			(*arr)[j] =
		}

		if i == 0 {
			next = (*arr)[index]
			(*arr)[index] = (*arr)[i]
		} else {
			(*arr)[index], next = next, (*arr)[index]
		}

	}
}

func BasicRotation(d int) {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	rotateMe(&a, d)
	fmt.Println(a)
}
