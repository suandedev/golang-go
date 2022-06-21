package main

import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"net/url"
)

var baseUrl = "htpp://localhost:8080"

type student struct {
	ID string
	Name string
	Grade int
}

func fetchUsers(ID string) ([]student, error) {
	var err error
	var client = &http.Client{}
	var data[]student

	var param = url.Values{}
	param.Set("id", ID)
	var playload = bytes.NewBufferString(param.Encode())

	request, err:= http.NewRequest("POST", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application-/x-www-from-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	var user1, err = fetchUsers("E001")
    if err != nil {
        fmt.Println("Error!", err.Error())
        return
    }

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}