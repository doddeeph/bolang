package main

import (
	"encoding/json"
	"fmt"
)

// https://dasarpemrogramangolang.novalagung.com/A-json.html
func main() {
	var strJson = `[
    	{"Name": "john wick", "Age": 27},
    	{"Name": "ethan hunt", "Age": 32}
	]`

	var data []user1

	var err = json.Unmarshal([]byte(strJson), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data[0], data[1])
}

type user1 struct {
	FullName string `json:"name"`
	Age      int
}
