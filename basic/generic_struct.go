package main

import "fmt"

// https://dasarpemrogramangolang.novalagung.com/A-golang-generics.html
func main() {
	var m1 = UserModel[int]{"noval", []int{1, 2, 3}}
	fmt.Println(m1)

	var m2 = UserModel[float64]{"noval", []float64{10, 11}}
	fmt.Println(m2)
}

func (m *UserModel[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

type UserModel[T int | float64] struct {
	Name   string
	Scores []T
}
