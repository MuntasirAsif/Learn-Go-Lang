package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Age      int      `json:"age"`
	Platform string   `json:"platform"`
	Password string   `json:"password"`
	Tags     []string `json:"tags"`
}

func main() {
	fmt.Println("Welcome to JSON ")
	// encodeJason()
	decodeJson()
}
func encodeJason() {
	lcoCourse := []course{
		{Name: "ReactJS Bootcamp", Age: 5, Platform: "Udemy", Password: "abc123", Tags: []string{"web-dev", "js"}},
		{Name: "Angular Bootcamp", Age: 4, Platform: "Udemy", Password: "abc123", Tags: []string{"app-dev", "flutter"}},
		{Name: "VueJS Bootcamp", Age: 3, Platform: "Udemy", Password: "abc123", Tags: []string{"backend-dev", "vuejs"}},
	}

	finalJson, _ := json.MarshalIndent(lcoCourse, "", "\t")
	fmt.Printf("%s", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(
		`{
                "course_name": "ReactJS Bootcamp",
                "age": 5,
                "platform": "Udemy",
                "password": "abc123",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
		`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// some cases where you just want to add data to key value pair
	var myOnlineData map[string]any
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}

}
