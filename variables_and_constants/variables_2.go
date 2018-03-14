package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	course := "Go Fundatmentals"

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)
	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {
	*course = "F# Fundamentals"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}
