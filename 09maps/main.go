package main

import "fmt"

func main() {
	fmt.Println("Maps in golang");

	languages := make(map[string]string);

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Length of map: ", len(languages))

	delete(languages, "RB")

	fmt.Println("List of all languages after delete: ", languages)

	fmt.Println("Js shorts for: ", languages["JS"])

	for key, value :=range languages{
		fmt.Println("Key:", key, "Value:", value)
	}
}