package main

import (
	"fmt"
)

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4])

	mySlice[1] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5] // 2 is inclusive, 5 is exclusive (so indexes 2, 3, 4 included)
	fmt.Println(sliceOfSlice)

	slice1 := mySlice[2:]
	fmt.Println(slice1)

	slice2 := mySlice[:5]
	fmt.Println(slice2)
}
