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

func GetName() string {
	url := "https://names.mcquay.me/api/v0"

	var name Name

	err := GetJson(url, &name)
	if err != nil {
			fmt.Printf("error getting name: %s\n", err.Error())
			return "Error, see terminal"
	}
	// else {
	//  return GetResponse(name.FirstName, name.LastName)
	// }
	 return GetResponse(name.FirstName, name.LastName)
	// return
}

func GetResponse(first string, last string) string {
	urlAugmented := "http://api.icndb.com/jokes/random?firstName=" + first + "&lastName=" + last + "&limitTo=nerdy"

	var response Response

	err := GetJson(urlAugmented, &response)
	if err != nil {
			fmt.Printf("error getting name: %s\n", err.Error())
			return "Error, see terminal"
	}
	// else {
		// fmt.Printf(response.Value.Joke)
		// return response.Value.Joke
	// }
		return response.Value.Joke
		// return

}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		client = &http.Client{Timeout: 10 * time.Second}
		var punchLine string = GetName()
		w.Write([]byte(punchLine))
	})
	http.ListenAndServe(":5000", nil)
}
