/*

Description:



*/

package main

import "fmt"

func ex6() {
	var n int

	fmt.Print("Enter a integer number: ")
	fmt.Scanln(&n)

	fmt.Println("The odd numbers between the chosen number and 1 are: ")
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
