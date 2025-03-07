/*

Description:

Ask the user to enter their age and whether they are over 18 or under 18 years old.

*/

package main

import "fmt"

func ex3() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	var isUnderage = age < 18
	var validAge = age >= 0

	if !validAge {
		fmt.Println("Invalid age! Please enter a non-negative number.")
	} else if isUnderage {
		fmt.Println("You are underage!")
	} else {
		fmt.Println("You are of legal age!")
	}
}
