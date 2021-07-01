package main

import (
	"belajar/thread"
	"fmt"
)

func main() {
	numberToTest := 1000 // number % 4 == 0 && !(number < 100)
	fmt.Println(thread.SumAlgA(numberToTest))
	fmt.Println(thread.SumAlgB(numberToTest))
	fmt.Println(thread.SumAlgC(numberToTest))
}
