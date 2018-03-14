package main

import (
	"fmt"
	"reflect"
)

// var (
// 	// name, course string
// 	// module	     float64
	
// 	// name, course, module = "Jason", "Docker Deep Dive", 3.2
// 	name = "Jason"
// 	course = "Docker Deep Dive"
// 	module = 3.2
// )

func main() {

	name := "Jason"
	// Won't compile if declared and not used at the function level.
	// course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("name is set to", name," and is of type ", reflect.TypeOf(name))
	fmt.Println("module is set to", module, "and is of type", reflect.TypeOf(name))
	fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr)
}