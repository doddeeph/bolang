package main

import (
	"fmt"
	"strings"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-fungsi-sebagai-parameter.html
*/
func main() {
	var data = []string{"wick", "jason", "ethan"}
	fmt.Println(data)

	var dataContainsO = filter(data, func(s string) bool {
		return strings.Contains(s, "o")
	})
	fmt.Println(dataContainsO)

	var dataLength5 = filterAlias(data, func(s string) bool {
		return len(s) == 5
	})
	fmt.Println(dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var res []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			res = append(res, each)
		}
	}
	return res
}

type FilterCallback func(string) bool

func filterAlias(data []string, callback FilterCallback) []string {
	var res []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			res = append(res, each)
		}
	}
	return res
}
