package main

import (
	"fmt"
	"net/url"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-url-parsing.html
*/
func main() {
	var strUrl = "http://kalipare.com:80/hello?name=john wick&age=27"
	fmt.Println("url:", strUrl)

	var u, e = url.Parse(strUrl)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println("protocol:", u.Scheme)
	fmt.Println("host:", u.Host)
	fmt.Println("path:", u.Path)

	var name = u.Query()["name"][0] // map[string][]string
	var age = u.Query()["age"][0]   // map[string][]string
	fmt.Println("name:", name, "age:", age)
}
