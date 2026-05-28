package main

import "fmt"

func main() {
	fmt.Println("Maps are key/value pairs or hash tables.")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("languages are: ", languages)

	// deleting by key
	delete(languages, "RB")

	fmt.Println("languages after deletion are: ", languages)

	// loop through Map
	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v\n", key, value)
	}

	// Prints map[] which is {}
	m := make(map[string]int)
	fmt.Println(m)

	// Prints map[] which is nil
	var m map[string]int
	fmt.Println(m)

}
