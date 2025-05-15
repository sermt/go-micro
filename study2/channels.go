package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func mainChannels() {
	unbufferedChannel := make(chan int)
	bufferedChannel := make(chan int, 10)

	// Unbuffered: se bloquea hasta que se recibe
	go func() {
		time.Sleep(1 * time.Second)
		unbufferedChannel <- rand.IntN(100)
	}()
	fmt.Println("Received from unbuffered channel:", <-unbufferedChannel)

	// Buffered: no se bloquea hasta que el buffer esté lleno
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		bufferedChannel <- rand.IntN(100)
	}
	close(bufferedChannel)

	for v := range bufferedChannel {
		fmt.Println("Received from buffered channel:", v)
	}
}
