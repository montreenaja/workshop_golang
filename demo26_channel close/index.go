package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}
	close(c)
}
func main() {
	ch := make(chan int, 5)
	go routine1(ch, 10)

	for true {
		value, ok := <-ch
		if !ok {
			fmt.Println("No more Data")
			break
		}
		fmt.Println(value)
	}

	//short form
	// for value := range ch {
	// 	fmt.Println(value)
	// }
	// fmt.Println("No more Data")

	time.Sleep(1 * time.Second)
}
