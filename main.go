package main

import (
	"belajar/thread"
	"fmt"
)

func main() {
	numberToTest := 1000 // applies multiply 10
	fmt.Println(thread.SumAlgA(numberToTest))
	fmt.Println(thread.SumAlgB(numberToTest))
	fmt.Println(thread.SumAlgC(numberToTest))
}
