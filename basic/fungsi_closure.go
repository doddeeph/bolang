package main

import "fmt"

/*
https://dasarpemrogramangolang.novalagung.com/A-fungsi-closure.html

Closure merupakan anonymous function atau fungsi tanpa nama
Bisa membuat fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi
Dimanfaatkan untuk membungkus suatu proses yang hanya dipakai sekali atau dipakai pada blok tertentu saja
*/
func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %v, max: %v\n", numbers, min, max)

	//Immediately-Invoked Function Expression (IIFE)
	min, max = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}(numbers)
	fmt.Printf("data: %v\nmin: %v, max: %v\n", numbers, min, max)

	var minNumber = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var filteredMinNumbers = func(min int) []int {
		var result []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			result = append(result, e)
		}
		return result
	}(minNumber)
	fmt.Println("numbers:", numbers)
	fmt.Printf("filtered by min %d --> numbers: %v\n", minNumber, filteredMinNumbers)

	//As return value
	var maxNumber = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, result = findMax(numbers, maxNumber)
	fmt.Println(howMany, result())
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
