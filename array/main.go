package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alamin-mahamud/go-data-structures/array/rotation"
)

func main() {
	d := flag.Int("d", 2, "rotate by this value")
	flag.Parse()
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	if *d > len(a) {
		os.Exit(1)
	}

	// rotation.SimpleRotation([]int{1, 2, 3, 4, 5, 6}, 3)
	// rotation.RotateOneByOne(&a, *d)
	rotation.JugglingRotation(&a, *d)
	fmt.Println(a)
}
