package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerfromPostRequest()
}

func PerfromPostRequest() {

	fmt.Println("post request in go")
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"courseName":"Learn Go Lang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	checkNilError(err)
	defer response.Body.Close()

	fmt.Printf("Response Type %T\n", response)
	fmt.Println("Response Status:", response.Status)
	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response Body:", string(content))

}

func PerformGetRequest() {
	fmt.Println("web request in go")
	const myUrl = "http://localhost:8000/"
	response, err := http.Get(myUrl)
	checkNilError(err)
	defer response.Body.Close()

	fmt.Printf("Response Type %T\n", response)
	fmt.Println("Response Status:", response.Status)

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response Body:", string(content))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
