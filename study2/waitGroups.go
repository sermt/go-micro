package main

import (
	"fmt"
	"sync"
)

func mainWaitGroups() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine 2")
	}()

	wg.Wait()
	fmt.Println("All goroutines have completed")
}
