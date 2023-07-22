package main

import (
	"fmt"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-map.html
*/
func main() {
	var data = map[string]int{}
	data["one"] = 1
	fmt.Println(data)

	var data1 = map[string]int{"januari": 50, "februari": 40}
	data1["maret"] = 34
	data1["april"] = 67
	fmt.Println(data1, len(data1))

	var data2 = make(map[string]int)
	data2["test"] = 123
	fmt.Println(data2)

	var data3 = *new(map[string]string)
	data3 = map[string]string{}
	data3["hello"] = "world"
	fmt.Println(data3)

	for key, value := range data1 {
		fmt.Println(key, "\t:", value)
	}

	delete(data1, "januari")
	fmt.Println(data1, len(data1))

	var key = "april"
	var value, exists = data1[key]
	if exists {
		fmt.Printf("Item %s: %d\n", key, value)
	} else {
		fmt.Printf("Item %s is not exists\n", key)
	}

	var data4 = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "female"},
		{"address": "no street", "id": "A123"},
		{"community": "chicken lovers"},
	}
	fmt.Println(data4)
}
