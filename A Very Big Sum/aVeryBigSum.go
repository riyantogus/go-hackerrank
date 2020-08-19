package main

import "fmt"

func main() {

	arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	avbs := aVeryBigSum(arr)

	fmt.Printf("%v\n", avbs)

}

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, each := range ar {
		sum += each
	}
	return sum
}
