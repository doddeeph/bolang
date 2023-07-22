package main

import "fmt"

/*
https://dasarpemrogramangolang.novalagung.com/A-struct.html

Go tidak memiliki class, konsep struct di golang mirip dengan konsep class pada OOP
*/
func main() {
	//Embedded Struct
	var s1 = student{person{"John", 10}, 4}
	fmt.Println("name:", s1.name, "grade:", s1.grade)

	var s2 = &s1
	fmt.Println(s1.name, s2.name)

	s2.name = "Ethan"
	fmt.Println(s1.name, s2.name)

	//Anonymous Struct
	var s3 = struct {
		person
		grade int
	}{person{"Wick", 10}, 4}
	fmt.Println(s3.name, s3.age, s3.grade)

	//Combine with slice
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	var allStudents1 = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	for _, student := range allStudents1 {
		fmt.Println(student)
	}
}

type person struct {
	name string
	age  int
}

type student struct {
	person
	grade int
}

//Nested struct adalah anonymous struct yang di-embed ke sebuah struct
//Teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup kompleks dengan proses decode hanya sekali.
//type student struct {
//	person struct {
//		name string
//		age  int
//	}
//	grade   int
//	hobbies []string
//}

//Tag property dalam struct
//Dimanfaatkan untuk keperluan encode/decode data json
//type person struct {
//	name string `tag1`
//	age  int    `tag2`
//}
