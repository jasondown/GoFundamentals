package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton", // Field names are optional if order of fields is known
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("Docker Deep Dive author:", DockerDeepDive.Author)
	DockerDeepDive.Rating = 4.5
	fmt.Println("Rating has changed from 5 to", DockerDeepDive.Rating)
}
