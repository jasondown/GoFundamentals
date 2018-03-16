package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("c:\\temp\\fileDoesNotExist.txt")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}
