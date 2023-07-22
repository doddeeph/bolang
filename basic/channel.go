package main

import (
	"fmt"
	"runtime"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-channel.html
*/
func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	//var sayHelloTo = func(who string) { //closure
	//	messages <- fmt.Sprintf("hello %s", who)
	//}
	//go sayHelloTo("john wick")
	//go sayHelloTo("ethan hunt")
	//go sayHelloTo("jason bourne")
	//var message1 = <-messages
	//var message2 = <-messages
	//var message3 = <-messages
	//fmt.Println(message1, "\n", message2, "\n", message3)

	var data = []string{"wick", "hunt", "bourne"}
	for _, e := range data {
		go func(who string) {
			messages <- fmt.Sprintf("hello %s", who)
		}(e)
	}
	for i := 0; i < len(data); i++ {
		printMsg1(messages)
	}
}

func printMsg1(what chan string) {
	fmt.Println(<-what)
}
