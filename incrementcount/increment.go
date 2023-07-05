package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func main() {
	counter := Counter{count: 0}
	fmt.Println(counter)
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.count) // Output: 1
}
