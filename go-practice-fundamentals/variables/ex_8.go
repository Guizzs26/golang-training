/*

Description:

Create a variable of type interface{} and assign it an integer.
Use type assertion to convert and print the value.

*/

package main

import "fmt"

func ex8() {
	var x interface{} = 42 // we can also use var x any = 42

	// Type assetion
	num := x.(int)

	fmt.Printf("O valor de x é: %d\n", num)
}
