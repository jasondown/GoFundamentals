package main

import (
	"fmt"
)

func main() {

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is: %v", key, value)
	}

	fmt.Println("\nC is:", testMap["C"])

	testMap["A"] = 100
	testMap["Jason"] = 39
	fmt.Println(testMap)

	fmt.Println("Deleting Jason...noooooooooo")
	delete(testMap, "Jason")
	fmt.Println(testMap)
}
