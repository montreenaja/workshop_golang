package main

import "fmt"

var count int = 5

func main() {
	fmt.Println("Hi")
	// explicit declar
	var test1 int = 10
	var test2 string = "man"
	var test3 bool = true

	//implicit delcar
	test4 := 0
	test5 := "Test implicit declar"

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(test5)
	run()

}

func run() {
	count += 5
	fmt.Println(count)
}
