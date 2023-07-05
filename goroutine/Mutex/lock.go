package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()

	// Critical section: modifying the shared resource
	counter++
}

func Lock() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
