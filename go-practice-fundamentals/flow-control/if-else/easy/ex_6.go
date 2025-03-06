/*

Description:

Ask the user to input three integers and determine which is the largest among them.

*/

package main

import "fmt"

func ex6() {
	var num1, num2, num3 int

	fmt.Println("Enter three integers numbers: ")
	fmt.Scanln(&num1, &num2, &num3)

	var max int = num1

	if num2 > max {
		max = num2
	}

	if num3 > max {
		max = num3
	}

	fmt.Printf("The largest number is: %d\n", max)
}
