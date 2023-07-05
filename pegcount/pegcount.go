package main

import (
	"fmt"
	"sync"
)

type PegCount struct {
	count int
	mu    sync.Mutex
}

func (p *PegCount) Increment() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.count++
}

func (p *PegCount) GetCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.count
}

func main() {
	pegCount := &PegCount{}

	// Simulate event occurrence
	pegCount.Increment()
	pegCount.Increment()
	pegCount.Increment()

	count := pegCount.GetCount()
	fmt.Printf("Peg count: %d\n", count)

	//...................................>

}
