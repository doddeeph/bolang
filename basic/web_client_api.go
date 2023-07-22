package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	for _, each := range users {
		fmt.Println(each)
	}

	var user1, err1 = fetchUser("E001")
	if err1 != nil {
		fmt.Println("Error!", err1.Error())
		return
	}

	fmt.Println(user1)
}

var baseUrl = "http://localhost:9009"

func fetchUsers() ([]mahasiswa1, error) {
	var err error
	var client = &http.Client{}
	var data []mahasiswa1

	req, err := http.NewRequest("GET", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (mahasiswa1, error) {
	var err error
	var client = &http.Client{}
	var data mahasiswa1

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	req, err := http.NewRequest("POST", baseUrl+"/user", payload)
	if err != nil {
		return data, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

type mahasiswa1 struct {
	ID    string
	Name  string
	Grade int
}
