/*

Description:

Ask the user to enter an integer and check if it is within the range of 10 to 50.

*/

package main

import "fmt"

func ex9() {
	var num int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&num)

	if num >= 10 && num <= 50 {
		fmt.Println("The number is within the range 10-50!")
	} else {
		fmt.Println("The number is OUT of the range 10-50!")
	}
}
