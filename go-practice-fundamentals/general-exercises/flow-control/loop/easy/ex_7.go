/*

Description:

Ask the user for a number N and calculate the factorial of that number.

*/

package main

import "fmt"

func ex7() {
	var n int
	fatorial := 1

	fmt.Println("Enter a integer number: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fatorial *= 1
	}

	fmt.Println("The fatorial of ", n, " is:", fatorial)
}
