package main

import (
	"fmt"
	"runtime"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-goroutine.html
*/
func main() {
	runtime.GOMAXPROCS(2)

	go printMsg(5, "Hello")
	printMsg(5, "How are you")

	var input string
	fmt.Scanln(&input)
}

func printMsg(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
