package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {

	fmt.Println("web request in go")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response Type %T\n", response)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	content := string(body)
	fmt.Println("Response Body:", content)
}
