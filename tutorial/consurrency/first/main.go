package main

import (
	"fmt"
	"sync"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func printSomethingWithWaitGroup(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	// sample1()
	// sample2()
	sample3()
}
func sample1() {
	go printSomething("This is the first thing to be printed!")
	time.Sleep(1 * time.Second)
	printSomething("This is the second thing to be printed!")
}

func sample2() {
	words := words()

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word))
	}

	time.Sleep(1 * time.Second)
	printSomething("This is the end of sample1")
}

func sample3() {
	var wg sync.WaitGroup

	words := words()

	wg.Add(len(words))

	for i, word := range words {
		go printSomethingWithWaitGroup(
			fmt.Sprintf("%d: %s", i, word),
			&wg)
	}

	wg.Wait()
	printSomething("This is the end of sample1")
}

func words() []string {
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	return words
}


