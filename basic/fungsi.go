package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-fungsi.html
*/
func main() {
	var names = []string{"John", "Wick"}
	printMessage("Hello", names)

	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	divideNumber(4, 0)
	divideNumber(8, -4)
}

func printMessage(message string, arr []string) {
	var joinArr = strings.Join(arr, " ")
	fmt.Println(message, joinArr)
}

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	var result = m / n
	fmt.Printf("%d / %d = %d\n", m, n, result)
}
