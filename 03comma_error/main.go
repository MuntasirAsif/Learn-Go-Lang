package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// input in go
	reader := bufio.NewReader(os.Stdin)

	/// comma ok idiom in go
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Your name is : ", input)
	fmt.Printf("The variable type is : %T\n", input)

	rating, _ := reader.ReadString('\n')
	fmt.Println("Your rating is : ", rating)
	fmt.Printf("The variable type is : %T\n", rating)
}