package main

import "fmt"

func variableInGo() {
	fmt.Println("This is a variable in Go:")
	var name string = "Muntasir"
	var age int = 25
	fmt.Println("Hello, " + name + "!")
	fmt.Printf("You are %d years old.\n", age)
	ageNext := 25
	fmt.Printf("Next year, you will be %d years old.\n", ageNext)
	fmt.Println("Goodbye, ", name, "!")
}
