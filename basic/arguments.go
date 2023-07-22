package main

import (
	"fmt"
	"os"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-command-line-args-flag.html
*/
func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)

	// go run basic/arguments.go banana potato "ice cream"
	// -> []string{"/var/folders/cr/rv5d85tx0q71s7dkj_3x0nqc0000gp/T/go-build850242537/b001/exe/arguments", "banana", "potato", "ice cream"}
	// -> []string{"banana", "potato", "ice cream"}
}
