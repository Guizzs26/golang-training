/*

Description:

Create a program that asks the user for a number N and displays the numbers from 1 to N
on the screen using a for loop.

*/

package main

import "fmt"

func ex1() {
	var n int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
