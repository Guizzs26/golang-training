/*

Description:

Ask the user for a number N and print all the even numbers from 1 to N.

*/

package main

import "fmt"

func ex5() {
	var n int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&n)

	fmt.Println("The even numbers between the chosen number and 1 are: ")
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
