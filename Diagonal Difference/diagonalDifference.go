package main

import (
	"fmt"
	"math"
)

func main() {

	ar := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	dDiff := diagonalDifference(ar)

	fmt.Printf("%d \n", dDiff)

}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var a, b int32
	n := len(arr)
	for i, each := range arr {
		a += each[i]
		b += each[n-i-1]
	}
	result := math.Abs(float64(a) - float64(b))
	return int32(result)
}
