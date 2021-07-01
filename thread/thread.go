package thread

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
	// mu := &sync.Mutex{}

	for i := 10; i <= numberToTest; i += 10 {
		var from, to int
		if i == 10 {
			from = 1
			to = 10
		} else {
			from = i - 9
			to = i
		}

		wg.Add(1)
		go func() {
			sum := 0
			for i := from; i <= to; i++ {
				sum += i
			}
			// mu.Lock()
			results += sum
			// mu.Unlock()
			wg.Done()
		}()
		wg.Wait()
	}

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
