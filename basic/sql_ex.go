package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// https://dasarpemrogramangolang.novalagung.com/A-sql.html
func main() {
	//sqlQuery()
	//sqlQueryRow()
	//sqlPrepare()
	//sqlExec()
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")

	// menggunakan metode prepared statement
	// stmt, err := db.Prepare("insert into tb_student values (?, ?, ?, ?)")
	// stmt.Exec("G001", "Galahad", 29, 2)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student1{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Println(result1)

	var result2 = student1{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Println(result2)

	var result3 = student1{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Println(result3)
}

// Query yang menghasilkan 1 baris record
func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student1{}
	var id = "E001"
	err = db.QueryRow("select name, grade from tb_student where id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student1

	for rows.Next() {
		var each = student1{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each)
	}
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=secret dbname=test sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, err
}

type student1 struct {
	id    string
	name  string
	age   int
	grade int
}
