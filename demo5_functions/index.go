package main

import "fmt"

func main() {
	fn1();
	fn2("Hello everyone!");
	fn3("i am batman", 7);
	const a,b = 2,3;
	fmt.Println(sum(a,b));
	fmt.Printf("%d+%d=%d\n",a,b, sum(a,b));
	x,y := swap(a,b);
	fmt.Printf("%d+%d => %d, %d\n",a,b, x,y);

	_x,_y, test1 := swap2(10,20);
	fmt.Printf("%d+%d => %d, %d, %s\n",a,b, _x,_y, test1);
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

func sum(a int, b int) int{
	return a + b;
}

func swap(a int, b int)(int, int){
	return b,a;
}

func swap2(a,b int)(x,y int, test1 string){
	y = a
	x = b
	test1 = "Result"
	return
}