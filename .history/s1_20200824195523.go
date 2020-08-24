package main

import (
	"fmt"
	// "log"
	"reflect"
)

func main() {
	// fmt.Println("Hello, world!!!")
	// println("Hello this is me")
	
	// variables and values
	// fmt.Println(1 + 2)
	// fmt.Println("this is my print", "and love to have more!")
	
	// in production e don't use fmt or geeky print(), we use log
	// log.Println("this will be logged")
	// -> there is no config!
	
	// MS name, other details, zap logger
	
	// easy or not? // Go is a kid of C
	// compare Go performance
	// 1- is the Golang faster than Java?
	  // it depends on: what we have on production
	// 2- arch
	  // java: JVM (isolated), java 11 to "just" compile to linux x86 
	  // golang: there is no no VM > you need to compile your code to x86 linux x86 windows
	// paradigms
	  // Java: OOP OOD
	  // Go: Not OOP, but Interfaces, Structs (something like classes) 
	// compare Clean Code (community stuff)
	  // java: cleaner
	  // golang: chain of conditional statements (passing errors)
	
	// varriables
	// name := "Massoud"
	// var family string = "Afrashteh" // java: string family
	// fmt.Println(name, family)
	// fmt.Printf("name: %s, family: %s", name, family)
	
	// constants
	// p := 3.14
	// const p float64 = 3.14
	// fmt.Println("p", p)
	// fmt.Println("type of p", reflect.ValueOf(p).Kind()) // python: type(p)
	
	var lessThanTen int = 5
	fmt.Println("lessThanTen", lessThanTen)
	
	checkTheMax := 6
	fmt.Println(checkTheMax, "type", reflect.ValueOf(checkTheMax).Kind())
	
	sum := lessThanTen + checkTheMax
	fmt.Println("sum",sum)
	
	// if
	if (sum == 11) {
		fmt.Println("sum is more than 11")
	} else {
		fmt.Println("sum is not more than 11")
	}
	
	if ("a" == "a") {
		fmt.Println("a is equal to a")
	}
	
	
	// exercise: write a program that can get the max (1000) print all of the odds and evens between 1-max
	
	
	// for errors
	if num := 10; num == 10 {
		fmt.Println("num is 10")
	}
	
	
	num := 23
	// for
	// java: for (int u = 0; u < 10; u++ {}
	for num < 23 {
		fmt.Println("num is less than 23")
	}
	
	for u := 0; u < 10; u++ {
		fmt.Println("num is less than 23")
	}
}
 
