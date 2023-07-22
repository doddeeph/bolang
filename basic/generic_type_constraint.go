package main

import "fmt"

// https://dasarpemrogramangolang.novalagung.com/A-golang-generics.html
func main() {
	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}
	fmt.Println("sum ints:", SumNumbers(ints))
	fmt.Println("sum floats:", SumNumbers(floats))
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

type Number interface {
	int64 | float64
}
