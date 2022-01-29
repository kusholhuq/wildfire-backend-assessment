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
		// fmt.Printf("First Name: %s\n", name.FirstName)
		// fmt.Printf("Last Name: %s\n", name.LastName)
		GetResponse(name.FirstName, name.LastName)
	}
}

func GetResponse(first string, last string) {
	// url := "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy"
	urlAugmented := "http://api.icndb.com/jokes/random?firstName=" + first + "&lastName=" + last + "&limitTo=nerdy"

	var response Response

	err := GetJson(urlAugmented, &response)
	if err != nil {
			fmt.Printf("error getting name: %s\n", err.Error())
	} else {
		// fmt.Printf("Type: %s\n", response.Type)
		// fmt.Printf("Value: %s\n", response.Value)
		// fmt.Printf("Value.Id: %s\n", response.Value.Id)
		fmt.Printf(response.Value.Joke)
		// fmt.Printf("Value.Categories: %s\n", response.Value.Categories)
		// fmt.Printf("PunchLine: %s\n", joke.PunchLine)
		// fmt.Printf("joke.Value.joke: %s\n", joke.Value["joke"])
	}
}


func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetName()
}
