package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("",numbers["two"])

	var colors = make(map[string]string)
	colors["green"] = "#0f0";
	colors["red"] = "#f00";
	colors["blue"] = "#00f";
	fmt.Println("",colors);
	fmt.Println("",colors["green"]);

	var courses = make(map[string]map[string]int)

	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 200
	courses["Android"]["price"] = 100
	courses["Android"]["code"] = 1234

	courses["iOs"] = make(map[string]int)
	courses["iOs"]["price"] = 100
	courses["iOs"]["code"] = 444
	fmt.Println(courses["iOs"]["code"]);
}