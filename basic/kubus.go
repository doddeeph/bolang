package main

import "math"

type Kubus1 struct {
	Sisi float64
}

func (k Kubus1) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

func (k Kubus1) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

func (k Kubus1) Keliling() float64 {
	return k.Sisi * 12
}
