package main

import (
	"fmt"
	"strings"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-interface-kosong.html
*/
func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
	var multiplied = secret.(float64) * 10 //casting to float64
	fmt.Println(secret, "--> multiplied by 10 is:", multiplied)

	secret = []string{"apple", "mango", "banana"}
	var fruits = strings.Join(secret.([]string), ",") //casting to slice of string
	fmt.Println(fruits, "are my favorite fruits")

	secret = &orang{"wick", 27}
	fmt.Println("name is:", secret.(*orang).name) //casting to object pointer

	var data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "mango", "banana"},
	}
	fmt.Println(data)

	data = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "mango", "banana"},
	}
	fmt.Println(data)

	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits1 = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits1 {
		fmt.Println(each)
	}
}

type orang struct {
	name string
	age  int
}
