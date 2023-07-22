package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg = calculate2(numbers...)
	//var avg = calculate2(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("rata-rata: %.2f", avg)
	fmt.Println(msg)

	var hobbies = []string{"eat", "sleep", "play"}
	myHobbies("John Wick", hobbies...)
}

func calculate2(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func myHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is %s, my hobies are: %s\n", name, hobbiesAsString)
}
