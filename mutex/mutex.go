package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.Mutex

func increment() {
	mutex.Lock()

	count++
	//defer mutex.Unlock()
}

func main() {
	// Multiple goroutines incrementing the count concurrently
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go increment()
	}
	fmt.Println("Count:", count) // Expected output: Count: 10
	// Wait for all goroutines to finish
	time.Sleep(time.Second)

}
