package main

import "fmt"

func main() {
	data := []string{"superman", "aquaman", "spiderman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("application running well")
				}
			}()
			panic("some error happen")
		}()
	}
}
