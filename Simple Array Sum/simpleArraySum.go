package main

import "fmt"

func main() {

	arr := []int32{1, 2, 3, 4, 10, 11}
	sas := simpleArraySum(arr)

	fmt.Printf("%d\n", sas)
}

func simpleArraySum(ar []int32) int32 {

	var sum int32
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}
