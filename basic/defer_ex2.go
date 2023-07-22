package main

import "fmt"

func main() {
	orderFood("burger")
	orderFood("pizza")
}

func orderFood(menu string) {
	defer fmt.Println("terima kasih, silahkan tunggu")
	if menu == "pizza" {
		fmt.Print("pilihan tepat!", " ")
		fmt.Print("pizza di tempat kami paling enak!", "\n")
		return
	}
	fmt.Println("pesanan anda:", menu)
	/*
		pesanan anda: burger
		terima kasih, silahkan tunggu
		pilihan tepat! pizza di tempat kami paling enak!
		terima kasih, silahkan tunggu
	*/
}
