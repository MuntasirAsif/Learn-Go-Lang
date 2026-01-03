package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	///conversion in go
	fmt.Println("Welcome to our shop")
	fmt.Println("Please rating our shop between 1 to 5")

	storeRating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", storeRating)

	numRating, err := strconv.ParseInt(strings.TrimSpace(storeRating), 10, 64)
	if err != nil {
		fmt.Println("Error in conversion", err)
	} else {
		fmt.Println("Added 1 to your rating ", numRating+1)
	}
}