package main

import (
	"fmt"

	concurrency "github.com/pobyzaarif/belajarGo/concurrency"
)

func main() {
	numberToTest := 1000 // number % 4 == 0 && !(number < 100)
	fmt.Println(concurrency.SumAlgA(numberToTest))
	fmt.Println(concurrency.SumAlgB(numberToTest))
	fmt.Println(concurrency.SumAlgC(numberToTest))
}
