package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer catch()

	var name string
	fmt.Print("type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
		//fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("application running well")
	}
}
