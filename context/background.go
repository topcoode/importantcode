package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Use the background context for some operation
	go performTask(ctx)

	// Wait for some time to allow the task to complete
	time.Sleep(2 * time.Second)
}

func performTask(ctx context.Context) {
	// Perform some task using the provided context
	fmt.Println("Task started", ctx)

	// Simulate task execution
	time.Sleep(1 * time.Second)

	// Check if the context is canceled
	if ctx.Err() != nil {
		fmt.Println("Task canceled")
		return
	}

	fmt.Println("Task completed", ctx)
}
