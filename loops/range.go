package main

import (
	"fmt"
)

func main() {

	coursesInProgress := []string{"Go Fundatmentals",
		"Making Your C# Code More Functional",
		"Working With Graph Algorithms in Python"}

	coursesCompleted := []string{"Making Your C# Code More Functional",
		"Introduction to Property-Based Testing With F#", "Applying Functional Principles in C#"}

	for _, i := range coursesInProgress {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nWoah now! We have course:-", i, "-in progress and completed. WAT!?!?")
			}
		}
	}

}
