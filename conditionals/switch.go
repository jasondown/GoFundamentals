package main

import (
	"fmt"
)

func main() {

	switch "docker" {
	case "linux":
		fmt.Println("\nRecommended Linux courses: ...")
	case "docker":
		fmt.Println("\nRecommended Docker courses: ...")
		fallthrough
	case "windows":
		fmt.Println("\nRecommended Windows courses: ...")
	default:
		fmt.Println("\nWe have no recommendations for you.")
	}
}
