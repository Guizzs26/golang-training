/*

Description:

Create a program that asks the user for a number N and calculates the sum of all numbers from 1 to N.

*/

package main

import "fmt"

func ex3() {
	var n, sum int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("The sum of the first", n, "numbers is:", sum)
}
