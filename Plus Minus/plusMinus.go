package main

import "fmt"

func main() {

	ar := []int32{-4, 3, -9, 0, 4, 1}

	plusMinus(ar)

}

func plusMinus(arr []int32) {
	var positive, negative, zero int32
	n := int32(len(arr))
	for _, each := range arr {
		if each > 0 {
			positive++
		} else if each < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n", float32(positive)/float32(n))
	fmt.Printf("%.6f\n", float32(negative)/float32(n))
	fmt.Printf("%.6f\n", float32(zero)/float32(n))
}
