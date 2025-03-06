/*

Description:

Ask the user to input two integers, and check if the first number is a multiple of
the second.

*/

package main

import "fmt"

func ex7() {
	var num1, num2 int

	fmt.Printf("Enter two integers numbers: ")
	fmt.Scanln(&num1, &num2)

	if num2 == 0 {
		fmt.Println("Division by zero is not allowed!")
		return
	}

	if num1%num2 == 0 {
		fmt.Println("The first number is a multiple of the second!")
	} else {
		fmt.Println("The first number is NOT a multiple of the second!")
	}
}
