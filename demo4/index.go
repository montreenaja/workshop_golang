package main

import "fmt"

func main() {
	fn1();
	fn2("Hello everyone!");
	fn3("i am batman", 7);
}

func fn1(){
	fmt.Println("Test1");
}
func fn2(msg string){
	fmt.Println(msg);
}
func fn3(title string , version int){
	fmt.Println(title);
	fmt.Println(version);
}