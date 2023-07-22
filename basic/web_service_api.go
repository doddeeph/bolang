package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getAllData)
	http.HandleFunc("/user", getSingleData)
	fmt.Println("starting web server at http://localhost:9009/")
	http.ListenAndServe(":9009", nil)

	// curl -X GET http://localhost:9009/users
	// [{"ID":"E001","Name":"ethan","Grade":21},{"ID":"W001","Name":"wick","Grade":22},{"ID":"B001","Name":"bourne","Grade":23},{"ID":"B002","Name":"bond","Grade":23}]

	// curl -X GET http://localhost:9009/user
	// User not found

	// curl -X POST http://localhost:9009/user -d "id=B002"
	// {"ID":"B002","Name":"bond","Grade":23}
}

func getAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func getSingleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		for _, each := range data {
			if each.ID == id {
				result, err := json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

var data = []mahasiswa{
	mahasiswa{"E001", "ethan", 21},
	mahasiswa{"W001", "wick", 22},
	mahasiswa{"B001", "bourne", 23},
	mahasiswa{"B002", "bond", 23},
}

type mahasiswa struct {
	ID    string
	Name  string
	Grade int
}
