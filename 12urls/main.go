package main

import (
	"fmt"
	"net/url"
)

const url1 = "https://jsonplaceholder.typicode.com:330/posts/1/?userId=1&postId=2"

func main() {
	// ...existing code...
	fmt.Println("urls in go")

	fmt.Println("URL is:", url1)

	result, _ := url.Parse(url1)
	fmt.Println("Parsed URL is:", result)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Scheme:", result.Scheme)


	qparams := result.Query()
	fmt.Printf("Query Params Type: %T\n", qparams)
	fmt.Println("Query Params:", qparams)
	for key, value := range qparams {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}