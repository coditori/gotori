package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

/*
challenge: modify this code so that the calls to updateMessage() on lines
28, 30, and 33 run as goroutines, and implement wait groups so that
the program runs properly, and prints out three different messages.
Then, write a test for all three function in this program: updateMessage(),
printSomething(), and main().
*/
func sample4() {
	msg = "Hello, world!"

	updateMessage("Hello, universe!")
	printMessage()

	updateMessage("Hello, cosmos!")
	printMessage()

	updateMessage("Hello, world!")
	printMessage()
}

var wg sync.WaitGroup

func updateMessage2(s string) {
	wg.Done()
	msg = s
}

func sample5SolutionForSample4() {
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage2("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage2("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage2("Hello, world!")
	wg.Wait()
	printMessage()
}
