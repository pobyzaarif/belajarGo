package main

import (
	"fmt"
	"sync"
)

var walletA = 1000
var walletB = 1000
var walletC = 1

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(2)
		go deductWalletA(&wg, 1)
		go deductWalletB(&wg, 1)
	}
	wg.Wait()

	go deductWalletC(1)

	fmt.Println(walletA)
	fmt.Println(walletB)
	fmt.Println(walletC)
}

var m = &sync.Mutex{}

func deductWalletA(wg *sync.WaitGroup, a int) {
	defer wg.Done()

	// prevent race condition lock this deduction proccess
	m.Lock()
	defer m.Unlock()

	walletA = walletA - a
}

func deductWalletB(wg *sync.WaitGroup, b int) {
	defer wg.Done()

	walletB = walletB - b
}

func deductWalletC(c int) {
	walletC = walletC - c
}
