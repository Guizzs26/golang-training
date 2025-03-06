/*

Description:

Create a program that determines whether a number is positive, negative, or zero.

*/

package main

import "fmt"

func ex4() {
	var num int

	fmt.Printf("Enter a integer number: ")
	fmt.Scanln(&num)

	var positiveNumber = num > 0
	var negativeNumber = num < 0

	if positiveNumber {
		fmt.Println("The number is a positive number.")
	} else if negativeNumber {
		fmt.Println("The number is a negative number.")
	} else {
		fmt.Println("The number is zero!")
	}
}
