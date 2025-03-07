/*

Description:

Create a program that prompts the user to input an integer and determines
whether it is even or odd.

*/

package main

import "fmt"

func ex1() {
	var num int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
