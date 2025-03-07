/*

Description:

Declare variables to store a name (string), age (int) and height (float64). Print this information in the format:

Name: John
Age: 25
Height: 1.75m

*/

package main

import "fmt"

func ex1() {
	// Declaring and initializing variables with short sintax
	name := "John"
	age := 25
	height := 1.75

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2fm\n", height)
}

/*

We can also declare and initialize variables this way:

var name string = "John"
var age int = 25
var height float64 = 1.75

*/
