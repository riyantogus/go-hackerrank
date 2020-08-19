package main

import "fmt"

func main() {

	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}

	compare := compareTriplets(a, b)

	fmt.Printf("%v\n", compare)
}

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {

	var alice, bob int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	return []int32{alice, bob}
}
