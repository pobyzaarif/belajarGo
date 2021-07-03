package concurrency

import (
	"sync"
)

func SumAlgA(numberToTest int) int {
	sum := 0
	for i := 1; i <= numberToTest; i++ {
		sum += i
	}

	return sum
}

func SumAlgB(numberToTest int) int {
	results := 0
	var wg sync.WaitGroup

	localSum := func(from, to int, wg *sync.WaitGroup) {
		defer wg.Done()
		sum := 0
		for i := from; i <= to; i++ {
			sum += i
		}
		results += sum
	}

	// This loop adding time consume

	// splitter := numberToTest / 5
	// for i := splitter; i <= numberToTest; i += splitter {
	// 	wg.Add(1)
	// 	if i == splitter {
	// 		go localSum(1, splitter, &wg)
	// 	} else {
	// 		go localSum(i-splitter+1, i, &wg)
	// 	}
	// 	wg.Wait()
	// }

	from1 := 1
	to1 := numberToTest / 4

	from2 := to1 + 1
	to2 := to1 + to1

	from3 := to2 + 1
	to3 := to2 + to1

	from4 := to3 + 1
	to4 := numberToTest

	wg.Add(4)
	go localSum(from1, to1, &wg)
	go localSum(from2, to2, &wg)
	go localSum(from3, to3, &wg)
	go localSum(from4, to4, &wg)
	wg.Wait()

	return results
}

func SumAlgC(numberToTest int) int {
	from1 := 1
	to1 := numberToTest / 2

	from2 := to1 + 1
	to2 := numberToTest

	c := make(chan int)
	localSum := func(from, to int, c chan int) {
		sum := 0
		for i := from; i <= to; i++ {
			sum += i
		}
		c <- sum
	}

	go localSum(from1, to1, c)
	go localSum(from2, to2, c)

	x, y := <-c, <-c

	return x + y
}
