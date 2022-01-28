package main

//Concurrency is the ability of a Large Application/program to run its Sub-programs or tasks independently and simultaneously.

//In Golang, Goroutines & Channels are used to make Concurrency easier.

import (
	"fmt"
	"time"
)

func say(word string) {
	for i := 0; i < 3; i++ {
		fmt.Println(word)
		time.Sleep(time.Microsecond * 100)
	}
}

func greeting(name string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- name
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func main() {
	//Using the goroutine to indicate Concurrency in go.
	start := time.Now()
	//We use the "go" keyword to set a goroutine

	for i := 0; i < 10; i++ {
		go say("Hello")
		go say("Nomso")
	}

	time.Sleep(time.Second)
	elapsed := time.Since(start)
	fmt.Printf(string(elapsed))

	//Using the Channel in Concurrency
	//Channels are defined using the "Make" function:
	//e.g VariableName := make(chan variableType)

	name := make(chan string)

	//the code below sends data to the channel
	go greeting("Hello", name)
}
