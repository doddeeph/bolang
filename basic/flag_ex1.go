package main

import (
	"flag"
	"fmt"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-command-line-args-flag.html
*/
func main() {
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 24, "type your age")
	var gender string
	flag.StringVar(&gender, "gender", "male", "type your gender")
	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
	fmt.Printf("gender\t: %s\n", gender)

	// go run basic/flag_ex1.go
	// name    : anonymous
	// age     : 24
	// gender  : male

	// go run basic/flag_ex1.go -name="john wick" -age=27  -gender=female
	// name    : john wick
	// age     : 27
	// gender  : female

	// go build basic/flag_ex1.go
	// ./flag_ex1 --help
	// Usage of ./flag_ex1:
	//   -age int
	//         type your age (default 24)
	//   -gender string
	//         type your gender (default "male")
	//   -name string
	//         type your name (default "anonymous")
}
