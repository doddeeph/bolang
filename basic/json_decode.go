package main

import (
	"encoding/json"
	"fmt"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-json.html
*/
func main() {
	var strJson = `{"name":"john wick","age":27}`
	var jsonData = []byte(strJson)

	var data user

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data.FullName, data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	fmt.Println(data1["name"], data1["age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodedData = data2.(map[string]interface{})
	fmt.Println(decodedData["name"], decodedData["age"])
}

type user struct {
	FullName string `json:"name"`
	Age      int
}
