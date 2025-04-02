package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		for {
			channel <- "ping"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			channel <- "pong"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// for { // infinito
	for i := 1; i <= 10; i++ { // limitado
		msg := <-channel
		fmt.Println(msg)
	}
}
