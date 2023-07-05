package main

import (
	"fmt"
	"sync"
)

func Goutine1(wg *sync.WaitGroup) {
	defer wg.Done()
	// Perform work for Goutine1
	fmt.Println("Goutine1")
}

func Goutine2(wg *sync.WaitGroup) {
	defer wg.Done()
	// Perform work for Goutine2
	fmt.Println("Goutine2")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go Goutine1(&wg)
	go Goutine2(&wg)

	wg.Wait()
	fmt.Println("All goroutines completed.")
}
