package main

import (
	"fmt"
	"strings"
)

/*
https://dasarpemrogramangolang.novalagung.com/A-method.html
*/
func main() {
	var s1 = siswa{"John Wick", 4}
	s1.sayHello()

	var firstName = s1.getNameAt(1)
	fmt.Println(firstName)

	var lastName = s1.getNameAt(2)
	fmt.Println(lastName)

	s1.changeName("Jason Bourne")
	fmt.Println(s1)

	s1.changeName1("Jason Bourne")
	fmt.Println(s1)

	var s2 = &siswa{"Ethan Hunt", 4}
	fmt.Println(s1.name, s2.name)
}

type siswa struct {
	name  string
	grade int
}

func (s siswa) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s siswa) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s siswa) changeName(name string) {
	s.name = name
}

func (s *siswa) changeName1(name string) {
	s.name = name
}
