package main

import (
	"fmt"
)

func main() {
	// functions
	// we can pass functions as arguments
	// don't forget to use variables names different than function ones

	// resources
	// https://gobyexample.com

	// slices: arrays with flexible size
	// we use "make" to create one like: arr := make([]string, 3)
	// when we need something fixed like weekdays: use array
	// when there is an need for change like a list of people: use slice

	// maps or dictionaries: they can have non-integer indexes
	// we use "make" to create one like: m := make(map[string]int)
	// they are pre sorted based on key content
	// sample: [day:2 friday:4 monday:0 z:5]

	m := make(map[string]string)
	m["name"] = "Ario"
	m["family"] = "Afrashteh"
	m["city"] = "Den Haag"
	fmt.Println(m)

	// range: iterate over collections (arrays, slices, maps, ...)

	var arr [12]string
	// arr := make([]string, 3)
	arr[0] = "Monday"
	// // 1 ""
	// // 2 ""

	// // arr[4] = "something"
	// fmt.Println("my slice len before append is", len(arr))

	// arr = append(arr, "something")

	// // fmt.Printf("this is my slice ")
	// fmt.Println("slice", arr)
	// fmt.Println("my slice len AFTER append is", len(arr))

	// arr := make([]string, 3)
	// res := fillArray(arr)
	// fmt.Println("my arr is", res)
	// fmt.Println("res array len is", len(res))

	// m := make(map[string]int)
	// m["monday"] = 0
	// m["friday"] = 4
	// m["a day"] = 2
	// m["z"] = 5
	// fmt.Println("my dict content", m)

	// delete(m, "friday")

	// fmt.Println("my dict after delete", m)

	// numberArr := []int{3, 10, 4, 20, 5, 1}
	// for i := 0; i < len(numberArr); i++ {
	// 	fmt.Printf("i: %d, value: %d\n", i, numberArr[i])
	// }

	// in java: you can just have values, IntStream
	// for (int v: arr) {

	// }

	// for k, v := range numberArr {
	// 	fmt.Printf("i: %d, value: %d\n", k, v)
	// }

	// for _, v := range numberArr {
	// 	fmt.Printf("value: %d\n", v)
	// }

	// return an error
	// for _, v := range 10 {
	// 	fmt.Printf("value: %d\n", v)
	// }

	// fmt.Println(myPrint("Massoud"))

	res := changeName(myPrint, "Masoud")
	fmt.Println("the result is", res)

}

func fillArray(arr []string) []string {
	arr[0] = "Monday"
	// 1 ""
	// 2 ""
	arr = append(arr, "something")
	arr = append(arr, "something else")
	return arr
}

func myPrint(s string) string {
	fmt.Println("myPrint is called")
	return s + " is changed!!!"
}

func changeName(f func(string) string, name string) string {
	fmt.Println("changeName is called")
	return f(name)
}
