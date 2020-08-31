package main

import (
	"fmt"
	"sort"
)

func main() {

	ar := []int32{1, 2, 3, 4, 5}

	miniMaxSum(ar)

}

func miniMaxSum(arr []int32) {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n := len(arr)
	var max, mini int64
	for i := 0; i < n-1; i++ {
		mini += int64(arr[i])
	}

	for i := n - 1; i > 0; i-- {
		max += int64(arr[i])
	}

	fmt.Println(mini, max)
}
