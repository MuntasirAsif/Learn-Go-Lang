package main

import "fmt"

func printValues[T string | int](value []T) {
	for _, v := range value {
		fmt.Println(v)
	}
}

type user[T string | int] struct {
	name T
	age  T
}

func main() {
	name := []string{"asif", "kodevio"}

	printValues(name)

	num := []int{1, 2, 3, 4, 5}

	printValues(num)

	user1 := []user[string]{
		{name: "ashif", age: "25"},
	}
	fmt.Println(user1)
}
