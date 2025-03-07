/*

Description:

Ask the user for a number X and display the multiplication table for that number from 1 to 10.

*/

package main

import "fmt"

func ex4() {
	var n int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&n)

	fmt.Println("The multiplication table for the number ", n, " is:")
	for i := 1; i <= 10; i++ {
		fmt.Println(n, " x ", i, " = ", n*i)
	}
}
