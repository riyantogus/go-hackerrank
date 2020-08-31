package main

import "fmt"

func main() {

	var n int32 = 6
	staircase(n)

}

func staircase(n int32) {

	var i, j, k int32

	for i = 0; i < n; i++ {

		for j = i; j < (n - 1); j++ {
			fmt.Printf(" ")
		}

		for k = 0; k < (i + 1); k++ {
			fmt.Printf("#")
		}

		fmt.Printf("\n")
	}

}
