/*

Description:

Write a Go program that checks whether three sides can form a triangle. For this to happen,
the sum of two sides must be greater than the third side.

*/

package main

import "fmt"

func ex12() {
	var a, b, c int

	fmt.Println("Enter the three sides of a triangle (integers a, b, c)")
	fmt.Scanln(&a, &b, &c)

	if a == 0 || b == 0 || c == 0 {
		fmt.Println("No valid triangle!")
	} else if a+b > c && a+c > b && b+c > a {
		fmt.Println("The sides form a triangle ")
	} else {
		fmt.Println("The sides do NOT form a triangle")
	}
}
