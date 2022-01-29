package main

import(
		"encoding/json"
		"net/http"
		"time"
		"fmt"
)

var client *http.Client

type Name struct {
	FirstName string `json:"first_name"`
	LastName string	`json:"last_name"`
}

type Response struct {
	Type string `json:"type"`
	Value ResponseValue `json:"value"`
}

type ResponseValue struct {
	Id int
	Joke string
	Categories []string
}

func GetName() {
	url := "https://names.mcquay.me/api/v0"

	var name Name

	err := GetJson(url, &name)
	if err != nil {
			fmt.Printf("error getting name: %s\n", err.Error())
	} else {
		fmt.Printf("First Name: %s\n", name.FirstName)
		fmt.Printf("Last Name: %s\n", name.LastName)
		GetResponse(name.FirstName, name.LastName)
	}
}
