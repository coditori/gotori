package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// IDEs
	// 1- vscode: geneeral IDE (Java, Golang, C#, ...)
	// 2- GoLand: 30 days free (specially written for Golang
	// 3- liteide: I like this
	
	// files
	// how to run golang programs?
	// go run filename.go
	
	// functions
	// in java: a function just have 1 and only one return, but it can accept any number of arguments
	// in golang: it can accept any number of arguments, but can have more than one return
	
	// arrays
	// in java: arrays are fixed sized new int[]{1,2,3} or new int[26], we can not print arrays with sout without Arrays.toString() 
	// in golang: it's the same

	// var name string = getName("Massoud") // my name is Massoud
	name := getName("Massoud")
	fmt.Println("returned value is", name)

	s, _ := sum(1, 5) // it should return sum, but if one of the variables are less than 0 it should return an error!
	fmt.Printf("sum result is %d\n", s)

	a, err := sum(1, -5) // it should return sum, but if one of the variables are less than 0 it should return an error!
	if err != nil {
		fmt.Printf("there was an error! %s\n", err)
	}
	fmt.Println(a)

	// default value for string is ""
	var arr [7]string
	arr[0] = "Monday"
	arr[6] = "Sunday"
	arr[4] = "Friday"
	// fmt.Println(arr[2])
	fmt.Println(arr)

	// default value for int is 0
	var intArr [10]int
	intArr[0] = 10
	intArr[5] = 50
	// fmt.Println(intArr)
	for i := 0; i < len(intArr); i++ {
		fmt.Printf("i: %d, value: %d\n", i, intArr[i])
	}

	// counting the occur
	// a, a, b, c, d > a2b1c1d1
	// a - ascii number of a
}

func sum(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("arguments should not be negative")
	}
	// error is an interface: interfaces in Go can be nil
	return a + b, nil
}

func getName(name string) string {
	return fmt.Sprintf("my name is %s", name)
}
