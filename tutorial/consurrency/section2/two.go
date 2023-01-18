package main

import (
	"fmt"
	"sync"
)

var msg2 string
var wg2 sync.WaitGroup

func updateMessage2(s string, m *sync.Mutex) {
	defer wg2.Done()
	m.Lock()
	msg2 = s
	m.Unlock()
}

func sample2() {
	msg2 = "Hello, world!"

	var mutex sync.Mutex

	wg2.Add(2)
	go updateMessage2("hello, universe!", &mutex)
	go updateMessage2("hello, cosmos!", &mutex)
	wg2.Wait()
	
	fmt.Println(msg2)
}