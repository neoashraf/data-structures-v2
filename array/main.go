package main

import (
	"flag"
	"fmt"

	"github.com/alamin-mahamud/go-data-structures/array/rotation"
)

func main() {
	// rotation.SimpleRotation([]int{1, 2, 3, 4, 5, 6}, 3)
	d := flag.Int("d", 2, "rotate by this value")
	flag.Parse()
	a := []int{1, 2, 3, 4, 5, 6}
	rotation.RotateOneByOne(&a, *d)
	fmt.Println(a)
}
