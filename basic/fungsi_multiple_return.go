package main

import (
	"fmt"
	"math"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-fungsi-multiple-return.html
*/
func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println(area, circumference)

	diameter = 30
	area, circumference = calculate1(diameter)
	fmt.Println(area, circumference)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

func calculate1(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}
