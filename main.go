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
