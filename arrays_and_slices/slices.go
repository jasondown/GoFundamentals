package main

import (
	"fmt"
)

func main() {

	myCourses := make([]string, 5, 10)
	myCourse2 := []string{"Go Fundamentals", "F# Jumpstart", "TDD with F#"}

	printSliceDetails(myCourses)
	printSliceDetails(myCourse2)
}

func printSliceDetails(slice []string) {
	fmt.Printf("\nLength is: %d.\nCapacity is: %d.", len(slice), cap(slice))
}
