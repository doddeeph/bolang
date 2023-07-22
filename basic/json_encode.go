package main

import (
	"encoding/json"
	"fmt"
)

// https://dasarpemrogramangolang.novalagung.com/A-json.html
func main() {
	var users = []user2{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var strJson = string(jsonData)
	fmt.Println(strJson)
}

type user2 struct {
	FullName string `json:"name"`
	Age      int
}
