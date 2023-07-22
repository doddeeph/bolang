package main

import "fmt"

/*
https://dasarpemrogramangolang.novalagung.com/A-pointer.html
*/
func main() {
	var numA = 4
	var numB = &numA
	fmt.Println(numA, &numA, numB, *numB)

	numA = 5
	fmt.Println(numA, &numA, numB, *numB)

	update(&numA, 10)
	fmt.Println(numA, &numA, numB, *numB)
}

func update(original *int, value int) {
	*original = value
}
