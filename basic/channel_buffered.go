package main

import (
	"fmt"
	"runtime"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-buffered-channel.html
*/
func main() {
	runtime.GOMAXPROCS(2)

	//Nilai buffered channel dimulai dari 0. Ketika nilainya adalah 3 berarti jumlah buffer maksimal ada 4.
	messages := make(chan int, 3)

	go func() {
		for {
			fmt.Println("receive data:", <-messages)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data:", i)
		messages <- i
	}

	//var input string
	//fmt.Scanln(&input)

	/*
		send data: 0
		send data: 1
		send data: 2
		send data: 3
		receive data: 0
		receive data: 1
		receive data: 2
		receive data: 3
		send data: 4
	*/
}
