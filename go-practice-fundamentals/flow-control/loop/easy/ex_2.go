/*

Description:

Create a program that asks the user for a number N and displays the numbers
from N to 1 in descending order.

*/

package main

import "fmt"

func ex2() {
	var n int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&n)

	fmt.Printf("Numbers from %d to 1 in descending order:\n", n)
	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}
