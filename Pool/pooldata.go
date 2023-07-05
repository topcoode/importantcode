package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Value int
}

func main() {
	// Create a new pool of Data objects
	dataPool := sync.Pool{
		New: func() interface{} {
			// Create a new Data object when the pool is empty
			return &Data{}
		},
	}

	// Get a Data object from the pool
	data := dataPool.Get().(*Data)

	// Set the value of the Data object
	data.Value = 42

	// Use the Data object
	fmt.Println("data-------->", data.Value)

	// Put the Data object back into the pool
	dataPool.Put(data)

	// Get another Data object from the pool (reusing the previously put object)
	data = dataPool.Get().(*Data)

	// Use the Data object again
	fmt.Println(data.Value)

	// Put the Data object back into the pool
	dataPool.Put(data)

}
