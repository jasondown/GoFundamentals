package main

import (
	"fmt"
)

func main() {

	const youSuck = "You suck! Consider another career"

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than second course.")
		if firstRank > 100 {
			fmt.Println(youSuck)
		}
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear... your first course must be shyte.")
		if secondRank > 100 {
			fmt.Println(youSuck)
		}
	} else {
		fmt.Println("\nEither both courses are doing the same or something shady is up.")
	}
}
