package main

import "fmt"

/*
https://dasarpemrogramangolang.novalagung.com/A-array.html
*/
func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits, len(fruits))

	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println(numbers, len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	fmt.Println(numbers1, len(numbers1))

	for i, fruit := range fruits {
		fmt.Printf("element %d: %s\n", i, fruit)
	}

	for _, num := range numbers {
		fmt.Printf("number: %d\n", num)
	}

	var fruits1 = make([]string, 2)
	fruits1[0] = "apple"
	fruits1[1] = "mango"
	fmt.Println(fruits1)
}
