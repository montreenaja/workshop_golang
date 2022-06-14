package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnWhileUseBreak()
}

func fnFor() {
	for index := 0; index <= 10; index++ {
		fmt.Printf("For Index %d\n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("While Index %d \n", index)
		index++
	}

}

func fnWhileUseBreak() {
	index := 0
	for true {
		if index > 5 {
			break;
		}
		fmt.Printf("While Break Index %d \n", index)
		index++
	}

}
