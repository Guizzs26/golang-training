/*

Description:

Ask the user to input two numbers and determine which is greater or if they are equal.

*/

package main

import "fmt"

func ex2() {
	var num1, num2 int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter a second number: ")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Println("The first number is greater than the second number!")
	} else if num2 > num1 {
		fmt.Println("The second number is greater than the first one!")
	} else {
		fmt.Println("They are equals!")
	}
}
